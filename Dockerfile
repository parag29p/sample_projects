FROM ubuntu:latest
LABEL "Author"="Parag"
LABEL "Project"="sample_go"
RUN apt update
RUN apt install apache2 -y
CMD ["/usr/sbin/apache2ctl", "-D", "FOREGROUND"]
EXPOSE 80 
WORKDIR /var/www/html
VOLUME /var/log/apache2
ADD HW-sample_v8 /var/www/html
