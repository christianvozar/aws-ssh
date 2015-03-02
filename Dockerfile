# Docker Definition for AWS SSH Inventory utility

FROM busybox:ubuntu-14.04
MAINTAINER Christian R. Vozar <christian@rogueethic.com>

ENV DEBIAN_FRONTEND=noninteractive

ADD cmd/aws-ssh/aws-ssh /aws-ssh
RUN chmod +x /aws-ssh

CMD ["/aws-ssh"]
