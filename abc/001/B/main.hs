main :: IO ()
calc (km)
  | km < 0.1 = "00"
  | km >= 0.1 && km <= 5 = "10"

main = do
  m <- fmap(\str -> read str :: Int) getLine
  km = (m `div` 1000)
