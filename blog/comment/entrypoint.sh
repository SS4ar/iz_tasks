#!/bin/sh
echo ${FLAG} >> /code/flag.txt

su -s /bin/sh -c 'flask run' user
