# MongoDB model
Feed documents: userId, city, category, listingId, title, createdAt, score.
Recommended compound indexes: {userId:1,createdAt:-1} and {city:1,category:1,createdAt:-1}.
Production topology: 3-node replica set with majority write concern for critical events.
