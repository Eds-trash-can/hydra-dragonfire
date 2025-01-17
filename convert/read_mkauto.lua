#!/usr/bin/env lua
require("spec")

local readers = {
	SliceByte = true,
	Byte = true,
	String = true,
	SliceField = true,
	Field = true,
	Bool = true,
	PointedThing = true,
}

local static_uses = {
	"[3]int16",
	"AOID"
}

local function generate(name)
	local fnname, index, child, childfn, childtype
	local type = name

	local open = name:find("%[")
	local clos = name:find("%]")

	if open == 1 then
		index = name:sub(open + 1, clos - 1)
		child = name:sub(clos + 1)
		childfn, childtype = generate(child)
		fnname = (index == "" and "Slice" or "Vec" .. index) .. childfn

		type = "[" .. index .. "]" .. childtype
	else
		fnname = camel_case(name)

		local c = name:sub(1, 1)
		if c == c:upper() then
			 type = "mt." .. name
		end
	end

	if not readers[fnname] then
		local fun = "func read" .. fnname .. "(l *lua.LState, val lua.LValue, ptr *" .. type  .. ") {\n"

		if child then
			fun = fun .. "\tif val.Type() != lua.LTTable {\n\t\tpanic(\"invalid value for "
				.. name .. ": must be a table\")\n\t}\n"

			if index == "" then
				fun = fun ..
[[
	tbl := val.(*lua.LTable)
	n := tbl.MaxN()
	*ptr = make(]] .. type .. [[, n)
	for i := range *ptr {
		read]] .. childfn .. [[(l, l.RawGetInt(tbl, i+1), &(*ptr)[i])
	}
]]
			else
				local n = tonumber(index)
				for i, v in ipairs({"x", "y", "z"}) do
					if i > n then
						break
					end

					fun = fun
						.. "\tread" .. childfn
						.. "(l, l.GetField(val, \"" .. v .. "\"), &(*ptr)[" .. (i - 1) .. "])\n"
				end
			end
		else
			fun = fun .. "\tif val.Type() != lua.LTNumber {\n\t\tpanic(\"invalid value for "
				.. name .. ": must be a number\")\n\t}\n"
				.. "\t*ptr = " .. type .. "(val.(lua.LNumber))\n"
		end

		fun = fun .. "}\n\n"

		readers[fnname] = fun
	end

	return fnname, type
end

for _, use in ipairs(static_uses) do
	generate(use)
end

local function signature(name, prefix, type)
	local camel = camel_case(name)
	return "func read" .. camel .. "(l *lua.LState, val lua.LValue, ptr *" .. prefix .. camel  .. ") {\n"
end

for name, fields in spairs(parse_spec("server/enum")) do
	local camel = camel_case(name)
	local fun = signature(name, "mt.")

	local impl = ""
	for _, var in ipairs(fields) do
		local equals = "*ptr = mt." .. apply_prefix(fields, var) .. "\n"

		if var == "no" then
			fun = fun .. "\tif val.Type() == lua.LTNil {\n\t\t" .. equals .. "\t\treturn\n\t}\n"
		else
			impl = impl .. "\tcase \"" .. var .. "\":\n\t\t" .. equals
		end
	end

	fun = fun
		.. "\tif val.Type() != lua.LTString {\n\t\tpanic(\"invalid value for "
		.. camel .. ": must be a string\")\n\t}\n"
		.. "\tstr := string(val.(lua.LString))\n"
		.. "\tswitch str {\n" .. impl
		.. "\tdefault:\n\t\tpanic(\"invalid value for " .. name .. ": \" + str)\n\t}\n}\n\n"

	readers[camel] = fun
end

for name, fields in spairs(parse_spec("server/flag")) do
	local camel = camel_case(name)
	local fun = signature(name, "mt.")
		.. "\tif val.Type() != lua.LTTable {\n\t\tpanic(\"invalid value for "
		.. camel .. ": must be a table\")\n\t}\n"

	for _, var in ipairs(fields) do
		fun = fun .. "\tif l.GetField(val, \"" .. var .. "\") == lua.LTrue {\n"
			.. "\t\t*ptr = *ptr | mt." .. apply_prefix(fields, var) .. "\n\t}\n"
	end

	fun = fun .. "}\n\n"
	readers[camel] = fun
end

local function fields_fromlua(fields, indent)
	local impl = ""

	for name, type in spairs(fields) do
		impl = impl .. indent .. "read" .. generate(type) .. "(l, l.GetField(val, \"" .. name .. "\"), &ptr."
			.. camel_case(name) .. ")\n"
	end

	return impl
end

for name, fields in spairs(parse_spec("server/struct", true)) do
	local camel = camel_case(name)
	readers[camel] = signature(name, "mt.")
		.. "\tif val.Type() != lua.LTTable {\n"
		.. "\t\tpanic(\"invalid value for " .. camel .. ": must be a table\")\n\t}\n"
		.. fields_fromlua(fields, "\t")
		.. "}\n\n"
end

local pkt_impl = ""

for name, fields in spairs(parse_spec("server/pkt", true)) do
	pkt_impl = pkt_impl
		.. "\tcase \"" .. name .. "\"" .. "" .. ":\n"
	 	.. "\t\tptr := &mt.ToSrv" .. camel_case(name) .. "{}\n"

	if next(fields) then
		pkt_impl = pkt_impl
			.. "\t\tval := l.CheckTable(3)\n"
			.. fields_fromlua(fields, "\t\t")
	end

	pkt_impl = pkt_impl
		.. "\t\treturn ptr\n"
end

local funcs = ""
for _, fn in spairs(readers) do
	if type(fn) == "string" then
		funcs = funcs .. fn
	end
end

local f = io.open("read_auto.go", "w")
f:write([[
// generated by read_mkauto.lua, DO NOT EDIT
package convert

import (
	"github.com/anon55555/mt"
	"github.com/yuin/gopher-lua"
)

]] .. funcs .. [[
func ReadCmd(l *lua.LState) mt.Cmd {
	str := l.CheckString(2)
	switch str {
]] .. pkt_impl .. [[
	}

	panic("invalid packet type: " + str)
}
]])
f:close()
