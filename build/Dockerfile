FROM registry.cn-hangzhou.aliyuncs.com/plugins/centos:7.9.2009
LABEL maintainers="Kube CustomResourceDefinitions Controller"
LABEL description="Kube CustomResourceDefinitions Controller"

COPY crd-controller /bin/crd-controller
RUN chmod +x /bin/crd-controller

ENTRYPOINT ["/bin/crd-controller"]
