local Hamming = {}

local function split(s)
    local result = {}
    for w in string.gmatch(s, ".") do
        table.insert(result, w)
    end
    return result
end

function Hamming.compute(a, b)
    local a_split = split(a)
    local b_split = split(b)

    if #a_split == #b_split then
        local distance = 0
        for k, v in pairs(a_split) do
            if b_split[k] ~= v then
                distance = distance + 1
            end
        end
        return distance
    else
        return -1
    end
end

return Hamming
