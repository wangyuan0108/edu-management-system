#!/bin/bash

cd /home/nginx/html/web
ls
echo "正在递归生成dist压缩包文件的MD5值"
md5sum dist.tgz > edu_system_rc2_md5.txt
echo "比较本地与远程MD5文件的内容"
diff edu_system_rc1_md5.txt edu_system_rc2_md5.txt
$result=$?
echo $result
if (($result == 0)); then
echo "两个文件的MD5相同"
else 
echo "两个文件的MD5不同, 退出"
exit
fi

echo "执行校验MD5文件流程"
bash dir.sh check_md5.sh
echo "解压dist文件"
tar -xzvf dist.tgz 
ls 
echo "删除dist压缩包"
rm -rf dist.tgz
echo "删除shell script"
rm -f dir.sh