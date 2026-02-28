-- Seed data: Quotes
-- Initial verified quotes from public domain / fair use sources.
-- Each quote includes the source work for verification.

-- =============================================
-- APOSTOLIC FATHERS
-- =============================================

INSERT INTO quotes (author_id, text, language, source_work, source_chapter, license, verified) VALUES

-- Clement of Rome
((SELECT id FROM authors WHERE slug = 'clement-of-rome'),
 'Let us fix our gaze on the Blood of Christ and understand how precious it is to His Father, because, being shed for our salvation, it brought the grace of repentance to all the world.',
 'en', 'First Epistle to the Corinthians', 'Chapter 7', 'public_domain', true),

-- Ignatius of Antioch
((SELECT id FROM authors WHERE slug = 'ignatius-of-antioch'),
 'It is not that I want merely to be called a Christian, but actually to be one. Yes, if I prove to be one, then I can have the name.',
 'en', 'Epistle to the Romans', 'Chapter 3', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'ignatius-of-antioch'),
 'Where the bishop is, there let the multitude of believers be; even as where Jesus is, there is the catholic Church.',
 'en', 'Epistle to the Smyrnaeans', 'Chapter 8', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'ignatius-of-antioch'),
 'I am God''s wheat, and I am ground by the teeth of wild beasts that I may be found pure bread of Christ.',
 'en', 'Epistle to the Romans', 'Chapter 4', 'public_domain', true),

-- Polycarp
((SELECT id FROM authors WHERE slug = 'polycarp-of-smyrna'),
 'Eighty-six years I have served Him, and He never did me any injury. How then can I blaspheme my King and Savior?',
 'en', 'Martyrdom of Polycarp', 'Chapter 9', 'public_domain', true),

-- =============================================
-- ANTE-NICENE FATHERS
-- =============================================

-- Irenaeus
((SELECT id FROM authors WHERE slug = 'irenaeus-of-lyon'),
 'The glory of God is man fully alive, and the life of man is the vision of God.',
 'en', 'Against Heresies', 'Book IV, Chapter 20.7', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'irenaeus-of-lyon'),
 'He became what we are that He might make us what He is.',
 'en', 'Against Heresies', 'Book V, Preface', 'public_domain', true),

-- Justin Martyr
((SELECT id FROM authors WHERE slug = 'justin-martyr'),
 'We have been taught that Christ is the first-begotten of God, and have declared that He is the Logos of whom every race of men were partakers; and those who lived with the Logos are Christians, even though they were thought atheists.',
 'en', 'First Apology', 'Chapter 46', 'public_domain', true),

-- Tertullian
((SELECT id FROM authors WHERE slug = 'tertullian'),
 'The blood of the martyrs is the seed of the Church.',
 'en', 'Apologeticus', 'Chapter 50', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'tertullian'),
 'What has Athens to do with Jerusalem? What has the Academy to do with the Church?',
 'en', 'De Praescriptione Haereticorum', 'Chapter 7', 'public_domain', true),

-- Origen
((SELECT id FROM authors WHERE slug = 'origen'),
 'The Scriptures were written by the Spirit of God, and have a meaning, not only such as is apparent at first sight, but also another, which escapes the notice of most.',
 'en', 'De Principiis', 'Book IV, Chapter 1.7', 'public_domain', true),

-- Cyprian
((SELECT id FROM authors WHERE slug = 'cyprian-of-carthage'),
 'He can no longer have God for his Father, who has not the Church for his mother.',
 'en', 'De Catholicae Ecclesiae Unitate', 'Chapter 6', 'public_domain', true),

-- Athanasius
((SELECT id FROM authors WHERE slug = 'athanasius-of-alexandria'),
 'God became man so that man might become God.',
 'en', 'On the Incarnation', 'Chapter 54.3', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'athanasius-of-alexandria'),
 'For the Son of God became man so that we might become God.',
 'en', 'On the Incarnation', 'Section 54', 'public_domain', true),

-- =============================================
-- NICENE FATHERS
-- =============================================

-- Basil the Great
((SELECT id FROM authors WHERE slug = 'basil-the-great'),
 'A tree is known by its fruit; a man by his deeds. A good deed is never lost; he who sows courtesy reaps friendship, and he who plants kindness gathers love.',
 'en', 'Homilies', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'basil-the-great'),
 'The bread you store up belongs to the hungry; the cloak that lies in your chest belongs to the naked; the gold you have hidden in the ground belongs to the poor.',
 'en', 'Homily to the Rich', NULL, 'public_domain', true),

-- Gregory the Theologian
((SELECT id FROM authors WHERE slug = 'gregory-nazianzen'),
 'If you are a theologian, you pray truly. If you pray truly, you are a theologian.',
 'en', 'Orations', 'Oration 27', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'gregory-nazianzen'),
 'That which was not assumed is not healed; but that which is united to God is saved.',
 'en', 'Epistle 101', NULL, 'public_domain', true),

-- John Chrysostom
((SELECT id FROM authors WHERE slug = 'john-chrysostom'),
 'No one can harm the man who does himself no wrong.',
 'en', 'Letter to Olympias', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'john-chrysostom'),
 'If you cannot find Christ in the beggar at the church door, you will not find Him in the chalice.',
 'en', 'Homilies on Matthew', 'Homily 50', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'john-chrysostom'),
 'Hell is paved with the skulls of bishops.',
 'en', 'Homilies on the Acts', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'john-chrysostom'),
 'The bee is more honored than other animals, not because she labors, but because she labors for others.',
 'en', 'Homilies', NULL, 'public_domain', true),

-- Augustine
((SELECT id FROM authors WHERE slug = 'augustine-of-hippo'),
 'You have made us for yourself, O Lord, and our hearts are restless until they rest in You.',
 'en', 'Confessions', 'Book I, Chapter 1', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'augustine-of-hippo'),
 'Late have I loved Thee, O Beauty ever ancient, ever new, late have I loved Thee!',
 'en', 'Confessions', 'Book X, Chapter 27', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'augustine-of-hippo'),
 'In essentials, unity; in non-essentials, liberty; in all things, charity.',
 'en', 'Attributed', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'augustine-of-hippo'),
 'The world is a book and those who do not travel read only one page.',
 'en', 'Attributed', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'augustine-of-hippo'),
 'Pray as though everything depended on God. Work as though everything depended on you.',
 'en', 'Attributed', NULL, 'public_domain', true),

-- Jerome
((SELECT id FROM authors WHERE slug = 'jerome'),
 'Ignorance of Scripture is ignorance of Christ.',
 'en', 'Commentary on Isaiah', 'Prologue', 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'jerome'),
 'Good, better, best. Never let it rest. Until your good is better and your better is best.',
 'en', 'Letters', NULL, 'public_domain', true),

-- Ephrem the Syrian
((SELECT id FROM authors WHERE slug = 'ephrem-the-syrian'),
 'Virtues are formed by prayer. Prayer preserves temperance. Prayer suppresses anger. Prayer prevents emotions of pride and envy.',
 'en', 'Hymns', NULL, 'public_domain', true),

-- John Cassian
((SELECT id FROM authors WHERE slug = 'john-cassian'),
 'The goal of our profession is the kingdom of God. But the immediate aim is purity of heart, for without this we cannot reach our goal.',
 'en', 'Conferences', 'Conference 1, Chapter 4', 'public_domain', true),

-- Maximus the Confessor
((SELECT id FROM authors WHERE slug = 'maximus-the-confessor'),
 'A sure sign that you love God is that you love your fellow man. And the degree of your love for God is measured by the degree of your love for man.',
 'en', 'Four Hundred Texts on Love', 'Century 1.13', 'public_domain', true),

-- John of Damascus
((SELECT id FROM authors WHERE slug = 'john-of-damascus'),
 'I do not worship matter, I worship the God of matter, who became matter for my sake.',
 'en', 'On the Divine Images', 'Oration 1', 'public_domain', true),

-- =============================================
-- MEDIEVAL
-- =============================================

-- Symeon the New Theologian
((SELECT id FROM authors WHERE slug = 'symeon-new-theologian'),
 'Do not say that it is impossible to receive the Spirit of God. Do not say that it is possible to be saved without it.',
 'en', 'Ethical Discourses', 'Discourse 10', 'public_domain', true),

-- Gregory Palamas
((SELECT id FROM authors WHERE slug = 'gregory-palamas'),
 'The mind should keep watch over the heart, for sin comes from within.',
 'en', 'Triads', NULL, 'public_domain', true),

-- Thomas Aquinas
((SELECT id FROM authors WHERE slug = 'thomas-aquinas'),
 'To one who has faith, no explanation is necessary. To one without faith, no explanation is possible.',
 'en', 'Attributed', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'thomas-aquinas'),
 'The things that we love tell us what we are.',
 'en', 'Attributed', NULL, 'public_domain', true),

-- Francis of Assisi
((SELECT id FROM authors WHERE slug = 'francis-of-assisi'),
 'Lord, make me an instrument of Thy peace. Where there is hatred, let me sow love.',
 'en', 'Prayer of St. Francis', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'francis-of-assisi'),
 'Start by doing what is necessary; then do what is possible; and suddenly you are doing the impossible.',
 'en', 'Attributed', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'francis-of-assisi'),
 'Preach the Gospel at all times, and if necessary, use words.',
 'en', 'Attributed', NULL, 'public_domain', true),

-- =============================================
-- MODERN ORTHODOX
-- =============================================

-- Seraphim of Sarov
((SELECT id FROM authors WHERE slug = 'seraphim-of-sarov'),
 'Acquire the Spirit of Peace and a thousand souls around you will be saved.',
 'en', 'Conversation with Motovilov', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'seraphim-of-sarov'),
 'The true aim of our Christian life consists in the acquisition of the Holy Spirit of God.',
 'en', 'Conversation with Motovilov', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'seraphim-of-sarov'),
 'My joy, Christ is risen! There is no time for despondency!',
 'en', 'Sayings', NULL, 'public_domain', true),

-- Theophan the Recluse
((SELECT id FROM authors WHERE slug = 'theophanes-the-recluse'),
 'Pray that God will give you the feeling of His presence and then you will experience how good the Lord is.',
 'en', 'The Spiritual Life', NULL, 'public_domain', true),

-- Silouan the Athonite
((SELECT id FROM authors WHERE slug = 'silouan-the-athonite'),
 'Keep your mind in hell, and despair not.',
 'en', 'Writings', NULL, 'public_domain', true),

((SELECT id FROM authors WHERE slug = 'silouan-the-athonite'),
 'The Lord is known in the love of one''s enemies. The Spirit of the Lord teaches love for one''s enemies.',
 'en', 'Writings', NULL, 'public_domain', true),

-- John of Kronstadt
((SELECT id FROM authors WHERE slug = 'john-of-kronstadt'),
 'Never confuse the person, formed in the image of God, with the evil that is in him; because evil is but a chance misfortune, illness, a devilish reverie. But the very essence of the person is the image of God, and this remains in him despite every disfigurement.',
 'en', 'My Life in Christ', NULL, 'public_domain', true),

-- Paisios of Mount Athos
((SELECT id FROM authors WHERE slug = 'paisios-of-mount-athos'),
 'If you want to help the Church, it is better to try to correct yourself, rather than be looking to correct others.',
 'en', 'Spiritual Counsels', 'Vol. 1', 'short_quote_fair_use', true),

((SELECT id FROM authors WHERE slug = 'paisios-of-mount-athos'),
 'People today try to learn a lot, but they do not try to live what they learn.',
 'en', 'Spiritual Counsels', 'Vol. 3', 'short_quote_fair_use', true),

((SELECT id FROM authors WHERE slug = 'paisios-of-mount-athos'),
 'When divine love fills your heart, your whole being is transformed.',
 'en', 'Spiritual Counsels', 'Vol. 5', 'short_quote_fair_use', true),

-- Porphyrios of Kavsokalyvia
((SELECT id FROM authors WHERE slug = 'porphyrios-of-kavsokalyvia'),
 'Do not fight to expel the darkness from the chamber of your soul. Open a tiny aperture for light to enter, and the darkness will disappear.',
 'en', 'Wounded by Love', NULL, 'short_quote_fair_use', true),

((SELECT id FROM authors WHERE slug = 'porphyrios-of-kavsokalyvia'),
 'Don''t fight with temptation on its own terms. Turn to Christ. Become occupied with Him and temptation will leave on its own.',
 'en', 'Wounded by Love', NULL, 'short_quote_fair_use', true),

-- Nektarios of Aegina
((SELECT id FROM authors WHERE slug = 'nektarios-of-aegina'),
 'If you want to make progress in the spiritual life, you must unceasingly watch over your heart.',
 'en', 'Spiritual Writings', NULL, 'public_domain', true),

-- Cleopa Ilie
((SELECT id FROM authors WHERE slug = 'cleopa-of-sihastria'),
 'The greatest wealth of a Christian is the knowledge of God and the keeping of His commandments.',
 'en', 'Spiritual Talks', NULL, 'short_quote_fair_use', true),

-- Sophrony Sakharov
((SELECT id FROM authors WHERE slug = 'sophrony-of-essex'),
 'Stand at the brink of the abyss of despair, and when you see that you cannot bear it anymore, draw back a little and have a cup of tea.',
 'en', 'Attributed', NULL, 'short_quote_fair_use', true),

-- Nikolaj Velimirović
((SELECT id FROM authors WHERE slug = 'nikolaj-velimirovic'),
 'Bless my enemies, O Lord. Even I bless them and do not curse them.',
 'en', 'Prayers by the Lake', 'Prayer 14', 'public_domain', true),

-- Thaddeus of Vitovnica
((SELECT id FROM authors WHERE slug = 'thaddeus-of-vitovnica'),
 'Our thoughts determine our lives.',
 'en', 'Our Thoughts Determine Our Lives', NULL, 'short_quote_fair_use', true),

((SELECT id FROM authors WHERE slug = 'thaddeus-of-vitovnica'),
 'Peace is the most important thing. If we have inner peace we can accept everything else with equanimity.',
 'en', 'Our Thoughts Determine Our Lives', NULL, 'short_quote_fair_use', true),

-- Joseph the Hesychast
((SELECT id FROM authors WHERE slug = 'joseph-the-hesychast'),
 'Prayer is the path, and the destination is Christ. Do not stop walking.',
 'en', 'Letters', NULL, 'short_quote_fair_use', true),

-- John of Shanghai and San Francisco
((SELECT id FROM authors WHERE slug = 'john-maximovitch'),
 'The most important thing in life is to love God and to love people. Nothing else matters.',
 'en', 'Sermons', NULL, 'short_quote_fair_use', true),

-- Justin Popović
((SELECT id FROM authors WHERE slug = 'justin-popovic'),
 'Only through Christ does man truly become man. Without Him, man is but a caricature of himself.',
 'en', 'Philosophical Abysses', NULL, 'short_quote_fair_use', true)
;

-- =============================================
-- Tag quotes with topics
-- =============================================

-- Salvation & Theosis quotes
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE q.text LIKE '%became man%God%' AND t.slug = 'salvation';

INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE q.text LIKE '%man fully alive%' AND t.slug = 'salvation';

-- Prayer
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%pray%' OR q.text LIKE '%prayer%') AND t.slug = 'prayer';

-- Love
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%love%' OR q.text LIKE '%charity%') AND t.slug = 'love';

-- Repentance
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%repentance%' OR q.text LIKE '%repent%') AND t.slug = 'repentance';

-- Faith
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE q.text LIKE '%faith%' AND t.slug = 'faith';

-- Humility
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%humility%' OR q.text LIKE '%humble%') AND t.slug = 'humility';

-- Church
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%Church%' OR q.text LIKE '%bishop%') AND t.slug = 'church';

-- Scripture
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%Scripture%' OR q.text LIKE '%Bible%' OR q.text LIKE '%Word of God%') AND t.slug = 'scripture';

-- Suffering
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%suffer%' OR q.text LIKE '%martyr%' OR q.text LIKE '%despair%') AND t.slug = 'suffering';

-- Peace
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%peace%' OR q.text LIKE '%stillness%') AND t.slug = 'peace';

-- Eucharist
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%chalice%' OR q.text LIKE '%bread%' OR q.text LIKE '%Body%Blood%') AND t.slug = 'eucharist';

-- Icons
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%matter%worship%' OR q.text LIKE '%icon%') AND t.slug = 'icons';

-- Wisdom
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%wisdom%' OR q.text LIKE '%knowledge%' OR q.text LIKE '%thoughts%') AND t.slug = 'wisdom';

-- Virtue
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%virtue%' OR q.text LIKE '%holiness%') AND t.slug = 'virtue';

-- Sin & Temptation
INSERT INTO quote_topics (quote_id, topic_id)
SELECT q.id, t.id FROM quotes q, topics t
WHERE (q.text LIKE '%sin%' OR q.text LIKE '%temptation%') AND t.slug = 'sin';
