For CentOS 5.5:

1. Install rsyslog if you don't already have it:

        $ yum install rsyslog

2. Use rsyslog instead of syslog:
3. Add [papertrail to rsyslog config](https://papertrailapp.com/systems/setup#0).
4. Stop `syslog` and start `rsyslog`:

        /etc/init.d/syslog stop && /etc/init.d/rsyslog start

5. Configure startups:

        $ /sbin/chkconfig syslog off
        $ /sbin/chkconfig --level 2345 rsyslog on

6. If SNMPd logging is too verbose: http://serverfault.com/questions/310640/reduce-snmpd-logging-verbosity
