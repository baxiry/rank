function grades(slc)  
    for k, v in pairs(slc) do
        if v > 37 then
            if (v+1)%5 == 0 then
                slc[k] = slc[k]+1
            end
            if (v+2)%5 == 0 then
                slc[k] = slc[k]+2
            end
        end
    end
    return slc
end

slc = {12,23,34,45,56,67,78,89,90}
grades(slc)
