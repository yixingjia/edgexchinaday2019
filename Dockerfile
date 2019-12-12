FROM photon:2.0

COPY htservice /
COPY src /src
COPY docker-entrypoint.sh /
COPY registData.json /
RUN chmod a+x htservice
RUN chmod a+x docker-entrypoint.sh
WORKDIR /
ENTRYPOINT ["bash","docker-entrypoint.sh"]
EXPOSE 9000
CMD ["./htservice"]
