FROM busybox:latest
LABEL maintainer="Jonathan Mainguy <jon@soh.re>"
ENV release=v0.2.0
ADD run.sh /usr/sbin/
RUN wget -O /usr/sbin/k8sCapcity_Linux_x86_64.tar.gz https://github.com/Jmainguy/k8sCapcity/releases/download/$release/k8sCapcity_0.2.0_linux_amd64.tar.gz

CMD ["/usr/sbin/run.sh"]
