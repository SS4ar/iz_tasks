# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: 'Star Wars' }, { name: 'Lord of the Rings' }])
#   Character.create(name: 'Luke', movie: movies.first)
u1 = User.create(username:'ffontaine', password:'|3!05|-|0(|<')
Secret.create(domain:'google.com',password:'@0ogfo!23AP42gK',user_id:'1')
Secret.create(domain:'jack.rayan',password:'wouldyoukindly',user_id:'1')

u2 = User.create(username:'afriden', password:'uu2V0mc06a')
Secret.create(domain:'amazon.se',password:'O2*OLs@1)dAm7',user_id:'2')

u3 = User.create(username:'vdarkholme', password:'p@ssw0rd')
Secret.create(domain:'hornet.com',password:'caG1Mih7hc&ui*',user_id:'3')
