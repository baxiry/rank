function howManyGames(price, diff, min, stoc)
    -- Return the number of games you can buy
    result = 0
    var = 0
    while stoc >= price do
        result = result + 1
        stoc  = stoc - price
        price = price-diff
        if price < min then
            price = min
        end
    end
    return result

end

r = howManyGames(20, 3, 6, 85)
print(r)
