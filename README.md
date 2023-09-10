# EnvoyCharts
A small data tracking/charting tool for envoy IQ solar systems

envoy-chart does two things
* runs in the background and polls the envoy every 10 minutes, putting that data into a database
* acts as a web server which displays the charts

# Still a work-in-progress
I'm just getting started.

Right now it supports a single graph (past 24-hours). I will work on adding 7-day and monthly graphs as well, after I collect enough data from my own system.

# Design decisions (or indecisions)
* 10 minute sampling frequency seems good enough. During initial testing I did every minute, but the graph isn't any more informative at this higher frequency, so I'll stick with 10 minutes for now. Eventually it will be configurable.
* go-echarts spits out Javascript Apache echarts (an older version at that). It seems nice enough for our purposes and is very simple to work with.
* sqlite because it works ; early experiments with objectbox were fun, but it was too fiddly for me
* go, because I like go

# To Do
* not have all the parameters hard coded, especially the hostname/ip address
* implement dns-sd in the go-envoy package so you don't have to set a hostname/ip at all
* more charts
* roll-up boxes (tables) for weekly/montly/yearly charts
