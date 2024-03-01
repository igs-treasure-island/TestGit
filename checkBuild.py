#!/usr/bin/python
#

import os
import sys
import mail

if len(sys.argv) < 2:
    print("Usage %s <jenkins batch job name>" % sys.argv[0])
    sys.exit()

JOB=sys.argv[1]

DIR = os.popen("getTriggerBuilds "+JOB).read()
FILENAME='Jenkins report at '+DIR
HTML=DIR+'.html'
ATT=DIR+'.zip'
output='No message'
att=None
if os.path.isfile(HTML):
    output=open(HTML,'r').read()
if os.path.isfile(ATT):
    att=[ATT]
mail.sendmail(FILENAME,output,att)
