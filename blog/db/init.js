db = db.getSiblingDB("blog");
db.createCollection('posts');
db.createCollection('users');

db.users.insertMany([
    {
        login: "admin",
        password: "h3g87uh487u43h8f7uh49u3"
    }
]);

db.posts.insertMany([
    {   "id": 0,
        "title":"Lunch",
        "src":"/images?file=post0.jpg", 
        "prevtext":"The capybara is like the celebrity of rodents, so chill...", 
        "content":"The capybara is like the celebrity of rodents, so chill that even other rodents forget how to properly chew nuts when they see it. They are so laid-back that in the most extreme situations, they can sit there looking as if they just had a cup of tea. The face of a capybara always seems to say, 'Oh, life, why are you so amusing!They even have their own interpersonal dynamics, where the alpha capybara is like a dear uncle always ready to listen to your problems and offer wise capybara-ish advice. In the evenings, all the capybaras gather for a pool party to discuss the latest rumors from the world of forest trails.The capybara is a character who, when it feels someone is watching, laughs inwardly, imagining itself as the star of the animal circus. And even if it looks like a big wet pillow, there's always a spark of joy in its eyes. The capybara is a true guru of rodent hedonism, the ultimate relaxer, and the charming idol of forest parties."
    },
    {   "id": 1,
        "title":"Friends",
        "src":"/images?file=post1.jpg", 
        "prevtext":"Capybara walks through the rodent scene like a true rockstar...", 
        "content":"Capybara walks through the rodent scene like a true rockstar, emitting an air of coolness that leaves even its fellow rodents struggling with their nut-cracking skills in its presence. Its laid-back attitude gives off the impression that it just finished leisurely sipping herbal tea, unfazed even in the most intense situations. The capybara's expression seems to eternally echo, Oh, life, why must you be so amusing!These creatures establish their own social hierarchy, with the alpha capybara assuming the role of a wise elder always ready to listen to your troubles and dispense sagely capybara-esque advice. As night falls, all the capybaras gather for a poolside celebration, exchanging the latest gossip from the world of wooded trails.The capybara is that character who, sensing eyes upon it, chuckles inwardly, imagining itself as the main act in the animal circus. Even when resembling a giant, damp cushion, there's a perpetual twinkle of joy in its eyes. The capybara stands as the true guru of rodent hedonism, the ultimate relaxation expert, and the charismatic idol of forest revelries."
    },
    {   "id": 2,
        "title":"Relax",
        "src":"/images?file=post2.jpg", 
        "prevtext":"The capybara, Hydrochoerus hydrochaeris, is the largest rodent in the world...", 
        "content":"The capybara, Hydrochoerus hydrochaeris, is the largest rodent in the world, native to South America. With a semi-aquatic lifestyle, these creatures are often found near bodies of water, such as rivers, lakes, and marshes. Capybaras have a robust and barrel-shaped body, webbed feet, and a blunt snout. They are highly social animals and live in groups, forming close-knit family units. Capybara communication involves vocalizations, such as barks and purrs, along with body language like grooming and nuzzling. Their gentle and docile nature makes them popular among both wildlife enthusiasts and locals. Capybaras are herbivores, primarily grazing on grasses and aquatic plants. Their efficient digestive systems allow them to extract nutrients from their plant-based diet. Due to their affinity for water, capybaras are excellent swimmers and can stay submerged for extended periods to evade predators. In various cultures, capybaras are regarded with a mix of curiosity and admiration. While their appearance may seem unusual, these creatures play a crucial role in the ecosystem by contributing to seed dispersal and controlling vegetation. Conservation efforts are in place to protect capybara populations, ensuring their continued presence in the diverse landscapes they inhabit."
    },
    {   "id": 3,
        "title":"Void",
        "src":"/images?file=post3.jpg", 
        "prevtext":"The capybara, scientifically known as Hydrochoerus hydrochaeris...", 
        "content":"The capybara, scientifically known as Hydrochoerus hydrochaeris, is a fascinating mammal native to South America. As the largest rodent on Earth, it boasts a unique and distinctive appearance. With a sturdy, barrel-shaped body, short legs, and partially webbed feet, the capybara is well-adapted to its semi-aquatic lifestyle.These social animals are often spotted near bodies of water, such as rivers, ponds, and marshes. Capybaras are known for their strong family bonds, living in groups that provide a sense of security and assistance in raising their young. Communication within capybara groups involves a range of vocalizations and subtle body language.Capybaras are herbivores, mainly feeding on grasses and aquatic plants. Their specialized digestive system enables them to efficiently extract nutrients from a plant-based diet. Given their affinity for water, capybaras are exceptional swimmers, capable of navigating through various aquatic environments.Beyond their biological significance, capybaras hold cultural intrigue. In different regions, they are admired for their gentle demeanor and unique features. Conservation initiatives are vital to ensure the protection of capybara populations, as they contribute to ecosystem balance through activities like seed dispersal and vegetation control. Appreciating the capybara's role in nature underscores the importance of preserving the rich biodiversity of South American landscapes."
    },
    {   "id": 4,
        "title":"Guard",
        "src":"/images?file=post4.jpg", 
        "prevtext":"The capybara, Hydrochoerus hydrochaeris, reigns as the largest rodent globally...", 
        "content":"The capybara, Hydrochoerus hydrochaeris, reigns as the largest rodent globally, and its presence is an intriguing facet of South American wildlife. Exhibiting a robust physique, short limbs, and partially webbed feet, the capybara is well-suited to its semi-aquatic habitat. Often found lounging near water bodies such as rivers and marshes, capybaras are highly social beings. They form tight-knit family groups, engaging in various forms of communication, from distinctive vocalizations to subtle body gestures. This communal lifestyle contributes to their survival and fosters a sense of camaraderie. As herbivores, capybaras graze on grasses and aquatic plants, showcasing a specialized digestive system that aids in extracting nutrients from their plant-centric diet. Their affinity for aquatic environments is further emphasized by their adept swimming skills, allowing them to navigate water with ease. Capybaras have piqued cultural interest, with their gentle nature and unique characteristics earning them admiration. Conservation efforts are imperative to safeguard these creatures, ensuring they continue to play a vital role in maintaining ecological balance. By appreciating the capybara's place in the intricate web of South American ecosystems, we underscore the importance of preserving biodiversity for generations to come."
    }
]);