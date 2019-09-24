FROM mysql:5.7
COPY entrypoint.sh /
COPY setup.sh /
COPY loadDatabase.sh /
COPY sql/ /sql
ENTRYPOINT ["bash", "/entrypoint.sh"]