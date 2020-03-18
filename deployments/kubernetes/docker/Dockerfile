FROM busybox:latest
LABEL maintainer="Jonathan Mainguy <jon@soh.re>"
ENV release=v0.2.3
RUN mkdir /opt
ADD run.sh /opt
WORKDIR /opt
RUN wget -O /opt/k8sCapcity_Linux_x86_64.tar.gz https://github.com/Jmainguy/k8sCapcity/releases/download/$release/k8sCapcity_Linux_x86_64.tar.gz
RUN chmod -R 777 /opt 
CMD ["/opt/run.sh"]
