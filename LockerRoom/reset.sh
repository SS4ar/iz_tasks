#!/bin/sh
cd /storage
bin/rails db:drop RAILS_ENV=development
bin/rails db:setup RAILS_ENV=development
bin/rails db:migrate RAILS_ENV=production
bundle exec rake assets:precompile RAILS_ENV=production
bin/rails s -b 0.0.0.0 -e production 
