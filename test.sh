D=`date`
BASEMSG="hello $D"
F=$FROM_EMAIL
T=$TO_EMAIL

echo "slack    $BASEMSG"      | ./msgr
echo "slack    $BASEMSG"      | ./msgr                 -code
echo "telegram $BASEMSG"      | ./msgr -where telegram
echo "telegram $BASEMSG"      | ./msgr -where telegram -code
echo "mailgun  $BASEMSG"      | ./msgr -where mailgun  -from "$F" -to "$T" -subject "mailgun $MSG"
echo "mailgun  $BASEMSG html" | ./msgr -where mailgun  -from "$F" -to "$T" -subject "mailgun $MSG html" -html
