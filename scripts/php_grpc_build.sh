#!/usr/bin/env bash
# 构建php的grpc相关依赖和扩展

# php版本
PHP_VERSION=""
# 操作系统发行版本号
OS_RELEASE_VERSION=""

GetPHPVersion() {
if [[ -x /usr/bin/php ]]; then
    $(/usr/bin/php -v | awk '/PHP\s([0-9])/ {print $2}' > /tmp/php_version.txt)
    PHP_VERSION=$(awk -F. '{print $1}' /tmp/php_version.txt)
    rm -rf /tmp/php_version.txt
elif [[ -x /usr/local/php/bin/php ]]; then
    $(/usr/local/php/bin/php -v | awk '/PHP\s([0-9])/ {print $2}' > /tmp/php_version.txt)
    PHP_VERSION=$(awk -F. '{print $1}' /tmp/php_version.txt)
    rm -rf /tmp/php_version.txt
else
    echo "can't found php"
    return
fi
}

GetOsVersion() {
    OS_RELEASE_VERSION=$(cat /etc/issue | awk '{print $1}')
    return
}


Main() {
    GetPHPVersion
    if [[ ${PHP_VERSION} -eq "" ]]; then
        exit -1
    elif [[ ${PHP_VERSION} -eq 5 ]]; then
        echo "php version: " ${PHP_VERSION}
    elif [[ ${PHP_VERSION} -eq 7 ]]; then
        echo "php version: " ${PHP_VERSION}
    fi
    GetOsVersion
    echo "os release version： " ${OS_RELEASE_VERSION}
}

Main