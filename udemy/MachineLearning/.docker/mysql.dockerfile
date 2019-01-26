FROM mysql:5.7

COPY ./.docker/mysqlsetup/my.cnf /etc/mysql/my.cnf
RUN deluser mysql
RUN useradd mysql
RUN chown -R mysql:mysql /var/lib/mysql
RUN chmod -R 777 /var/lib/mysql