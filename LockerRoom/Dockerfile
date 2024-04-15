FROM ruby:2.5.5-alpine

WORKDIR /storage

COPY . /storage/

RUN apk update \
    && apk add build-base sqlite-dev shared-mime-info nodejs

RUN gem install bundler -v 2.3.27

#CMD ["gem", "list", "bundler"]
RUN bundle install

CMD ["sh", "reset.sh"]
