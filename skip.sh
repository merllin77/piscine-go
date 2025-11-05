#ls -l | sed -n 'p;n'
#ls -l --quoting-style=shell-escape | awk 'NR % 2 == 1'
ls -l --quoting-style=shell-escape | sed -n 'n;p'