FROM scratch

ADD openai-proxy /openai-proxy

EXPOSE 8890

ENTRYPOINT ["/openai-proxy"]
