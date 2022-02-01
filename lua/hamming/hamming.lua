local Hamming = {}

function Hamming.compute(a, b)

end

local function splitString(s, delimiter)
    result = {}
    for match in (s .. delimiter):gmatch("" .. delimiter) do
        table.insert(result, match)
    end
    return result
end

return Hamming
