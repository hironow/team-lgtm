FROM golang:1.12

COPY run_tests.sh /run_tests.sh
RUN ["chmod", "+x", "/run_tests.sh"]
ENTRYPOINT ["/run_tests.sh"]
