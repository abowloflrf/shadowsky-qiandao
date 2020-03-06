# shadowsky 签到

![Go](https://github.com/abowloflrf/shadowsky-qiandao/workflows/Go/badge.svg?branch=master)

```shell
git clone https://github.com/abowloflrf/shadowsky-qiandao.git
cd shadowsky-qiandao
go build .
cp .env.example .env # remember to edit the config file
./shadowsky-qiandao > job.log 2>&1 & # run cronjob in background and redirect stdout log to job.log file
```
