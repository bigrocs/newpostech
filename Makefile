.PHONY: git
git:
	git add .
	git commit -m"自动提交 git 代码"
	git push
tag:
	git push --tags
pem:
	openssl pkcs12 -in ./testKey/1039702660-商户.pfx -nodes -out ./testKey/1039702660-商户.pem
cer:
	openssl x509 -in ./testKey/1039683724-联盟.cer -pubkey  -noout > ./testKey/1039683724-联盟.pem
key:
	openssl pkcs12 -in ./testKey/1039702660-商户.pfx -nocerts -nodes -out ./testKey/1039702660-商户.key
keyPri:
	openssl rsa -in ./testKey/1039702660-商户.key -out ./testKey/1039702660-商户_pri.key
keyPub:
	openssl rsa -in ./testKey/1039702660-商户.key -pubout -out ./testKey/1039702660-商户_pub.key