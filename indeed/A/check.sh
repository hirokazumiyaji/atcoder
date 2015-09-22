for f in `ls *.txt`; do
  echo $f
  cat $f | python main.py
done
