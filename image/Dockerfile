FROM ubuntu

RUN set -x && \
  apt-get update && \
  apt-get install -y rsyslog rsyslog-gnutls

RUN set -x && \
  touch /var/log/rsyslog.log && \
  ln -sf /dev/stdout /var/log/rsyslog.log

RUN mkdir /etc/rsyslog_certs
ADD ./ca.pem /etc/rsyslog_certs/ca.pem
ADD ./cert.pem /etc/rsyslog_certs/cert.pem
ADD ./key.pem /etc/rsyslog_certs/key.pem

ADD ./rsyslog.conf /etc/rsyslog.conf

CMD [ "rsyslogd", "-n" ]
