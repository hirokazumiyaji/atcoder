main :: IO ()
main = do
  h1 <- fmap(\str -> read str :: Int) getLine
  h2 <- fmap(\str -> read str :: Int) getLine
  print $ h1 - h2
