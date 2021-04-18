cpu=$(top -b -n1 | grep "xxx" | head -1 | awk '{print $9}')
pid=$(top -b -n1 | grep "xxx" | head -1 | awk '{print $1}')
echo $cpu
echo $pid
cpu1=$(top -b -n1 | grep "___go_build_" | head -1 | awk '{print $9}')
pid1=$(top -b -n1 | grep "___go_build_" | head -1 | awk '{print $1}')
echo $cpu1
echo $pid1
if [ $cpu > 85 ]
kill -9 $pid
then echo "kill -9 $pid";
fi