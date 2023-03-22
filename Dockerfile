FROM amazonlinux:2
LABEL maintainer="marcel"
LABEL server="EC2 Magento"

EXPOSE 80 443

RUN amazon-linux-extras install nginx1

RUN yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
RUN yum -y install https://rpms.remirepo
RUN yum -y install yum-utils rpl
RUN yum-config-manager --enable remi-php82
RUN rpl [remi-php82] "[remi-php82]\npriority=1" /etc/yum.repos.d/remi-php82.repo -e

# install deps
RUN yum install -y
RUN pecl channel-update pecl.php.net
RUN pecl install redis imagick

RUN wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.17.8-x86_64.rpm
RUN wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.17.8-x86_64.rpm.sha512
RUN sha512sum -c elasticsearch-7.17.8-x86_64.rpm.sha512
RUN rpm --install elasticsearch-7.17.8-x86_64.rpm