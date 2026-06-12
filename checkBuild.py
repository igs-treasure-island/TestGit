#!/usr/bin/python
#

import os
import sys
import mail

if len(sys.argv) < 2:
    print("Usage %s <jenkins batch job name>" % sys.argv[0])
    sys.exit()

DIR = 'Foo'
FILENAME='Jenkins report at '+DIR
output='No message'
att=None
mail.sendmail(FILENAME,output,att)
