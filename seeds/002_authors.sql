-- Seed data: Authors
-- Tier 1: Apostolic Fathers
-- Tier 2: Ante-Nicene Fathers
-- Tier 3: Nicene & Post-Nicene Fathers (the "Golden Age")
-- Tier 4: Medieval theologians
-- Tier 5: Reformation-era and later
-- Tier 6: Modern Orthodox Saints & Elders

-- =============================================
-- TIER 1: APOSTOLIC FATHERS (c. 50-150 AD)
-- =============================================

INSERT INTO authors (slug, name, title, born_year, died_year, era, tradition, bio_short, canonized, copyright_status, wikipedia_url, wikimedia_category) VALUES
('clement-of-rome', 'Clement of Rome', 'Bishop of Rome', 35, 99, 'apostolic', 'pre_schism',
 'Third successor of Peter as Bishop of Rome. Author of the Epistle to the Corinthians, one of the earliest non-canonical Christian writings.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Pope_Clement_I', 'Pope Clement I'),

('ignatius-of-antioch', 'Ignatius of Antioch', 'Bishop of Antioch, Martyr', 35, 108, 'apostolic', 'pre_schism',
 'Student of the Apostle John. Wrote seven epistles while being transported to Rome for martyrdom. Key witness to early Church structure.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Ignatius_of_Antioch', 'Ignatius of Antioch'),

('polycarp-of-smyrna', 'Polycarp of Smyrna', 'Bishop of Smyrna, Martyr', 69, 155, 'apostolic', 'pre_schism',
 'Disciple of the Apostle John. His martyrdom account is one of the earliest recorded. Said "Eighty-six years I have served Him" before his death.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Polycarp', 'Polycarp'),

('didache', 'The Didache', NULL, NULL, NULL, 'apostolic', 'pre_schism',
 'Anonymous first-century teaching manual, also called "Teaching of the Twelve Apostles." Earliest known catechism.',
 false, 'public_domain', 'https://en.wikipedia.org/wiki/Didache', NULL),

-- =============================================
-- TIER 2: ANTE-NICENE FATHERS (c. 100-325 AD)
-- =============================================

('irenaeus-of-lyon', 'Irenaeus of Lyon', 'Bishop of Lyon', 130, 202, 'ante_nicene', 'pre_schism',
 'Student of Polycarp. His "Against Heresies" is a masterwork defending orthodox Christianity against Gnosticism.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Irenaeus', 'Irenaeus'),

('justin-martyr', 'Justin Martyr', 'Philosopher, Martyr', 100, 165, 'ante_nicene', 'pre_schism',
 'First great Christian apologist. Sought truth through philosophy before finding it in Christ. Martyred in Rome.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Justin_Martyr', 'Justin Martyr'),

('tertullian', 'Tertullian', 'Presbyter of Carthage', 155, 220, 'ante_nicene', 'pre_schism',
 'Father of Latin Christianity. Coined the term "Trinity" (trinitas). Prolific writer and fierce defender of the faith.',
 false, 'public_domain', 'https://en.wikipedia.org/wiki/Tertullian', 'Tertullian'),

('origen', 'Origen of Alexandria', 'Head of Catechetical School', 185, 253, 'ante_nicene', 'pre_schism',
 'One of the most prolific writers of early Christianity. Pioneer of biblical scholarship and systematic theology.',
 false, 'public_domain', 'https://en.wikipedia.org/wiki/Origen', 'Origen'),

('cyprian-of-carthage', 'Cyprian of Carthage', 'Bishop of Carthage, Martyr', 210, 258, 'ante_nicene', 'pre_schism',
 'Bishop during persecutions. His "On the Unity of the Church" remains a foundational ecclesiological text.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Cyprian', 'Cyprian'),

('athanasius-of-alexandria', 'Athanasius of Alexandria', 'Archbishop of Alexandria', 296, 373, 'ante_nicene', 'pre_schism',
 'Champion of Trinitarian orthodoxy against Arianism. Exiled five times for defending the Nicene faith. "Athanasius contra mundum."',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Athanasius_of_Alexandria', 'Athanasius of Alexandria'),

-- =============================================
-- TIER 3: NICENE & POST-NICENE FATHERS (c. 325-800)
-- =============================================

('basil-the-great', 'Basil the Great', 'Archbishop of Caesarea', 330, 379, 'nicene', 'pre_schism',
 'One of the Cappadocian Fathers. Founded Eastern monasticism. His liturgy is still celebrated in the Orthodox Church.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Basil_of_Caesarea', 'Basil of Caesarea'),

('gregory-nazianzen', 'Gregory the Theologian', 'Archbishop of Constantinople', 329, 390, 'nicene', 'pre_schism',
 'Cappadocian Father. Called "The Theologian" for his profound orations on the Trinity. Brief Patriarch of Constantinople.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Gregory_of_Nazianzus', 'Gregory of Nazianzus'),

('gregory-of-nyssa', 'Gregory of Nyssa', 'Bishop of Nyssa', 335, 395, 'nicene', 'pre_schism',
 'Cappadocian Father. Mystic theologian and brother of Basil. Pioneer of apophatic theology and the concept of epektasis.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Gregory_of_Nyssa', 'Gregory of Nyssa'),

('john-chrysostom', 'John Chrysostom', 'Archbishop of Constantinople', 349, 407, 'nicene', 'pre_schism',
 '"Golden Mouth" — greatest preacher of the early Church. His homilies remain among the most read patristic works.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/John_Chrysostom', 'John Chrysostom'),

('augustine-of-hippo', 'Augustine of Hippo', 'Bishop of Hippo', 354, 430, 'nicene', 'pre_schism',
 'Most influential Western Father. His "Confessions" and "City of God" shaped Western theology, philosophy, and literature.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Augustine_of_Hippo', 'Augustine of Hippo'),

('jerome', 'Jerome', 'Priest, Doctor of the Church', 347, 420, 'nicene', 'pre_schism',
 'Translator of the Vulgate Bible. Most learned of the Latin Fathers. Patron saint of translators and scholars.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Jerome', 'Jerome'),

('cyril-of-alexandria', 'Cyril of Alexandria', 'Patriarch of Alexandria', 376, 444, 'nicene', 'pre_schism',
 'Central figure at the Council of Ephesus. Defender of the title Theotokos and the unity of Christ''s person.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Cyril_of_Alexandria', 'Cyril of Alexandria'),

('maximus-the-confessor', 'Maximus the Confessor', 'Monk, Confessor', 580, 662, 'post_nicene', 'pre_schism',
 'Defender of Christ''s two wills against Monothelitism. His tongue and hand were cut off for his confession. Profound mystic theologian.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Maximus_the_Confessor', 'Maximus the Confessor'),

('john-of-damascus', 'John of Damascus', 'Monk, Doctor of the Church', 676, 749, 'post_nicene', 'pre_schism',
 'Last of the great Eastern Fathers. Defender of icons. His "Exact Exposition of the Orthodox Faith" is a systematic masterpiece.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/John_of_Damascus', 'John of Damascus'),

('ephrem-the-syrian', 'Ephrem the Syrian', 'Deacon of Edessa', 306, 373, 'nicene', 'pre_schism',
 'The "Harp of the Spirit." Greatest of the Syrian Church Fathers. Wrote theological poetry and hymns of extraordinary beauty.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Ephrem_the_Syrian', 'Ephrem the Syrian'),

('john-cassian', 'John Cassian', 'Monk, Priest', 360, 435, 'nicene', 'pre_schism',
 'Brought Eastern monasticism to the West. His "Conferences" and "Institutes" are foundational monastic texts.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/John_Cassian', 'John Cassian'),

('desert-fathers', 'The Desert Fathers', NULL, NULL, NULL, 'nicene', 'pre_schism',
 'Anonymous and named hermits, ascetics, and monks living in the Egyptian desert from the 3rd century. Source of the "Sayings of the Desert Fathers."',
 false, 'public_domain', 'https://en.wikipedia.org/wiki/Desert_Fathers', NULL),

-- =============================================
-- TIER 4: MEDIEVAL (c. 800-1500)
-- =============================================

('symeon-new-theologian', 'Symeon the New Theologian', 'Abbot', 949, 1022, 'medieval', 'orthodox',
 'One of only three saints granted the title "Theologian" in Orthodoxy. Mystic who spoke of direct experience of divine light.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Symeon_the_New_Theologian', 'Symeon the New Theologian'),

('gregory-palamas', 'Gregory Palamas', 'Archbishop of Thessalonica', 1296, 1359, 'medieval', 'orthodox',
 'Champion of Hesychasm. Defended the reality of the uncreated divine light. Articulated the essence-energies distinction.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Gregory_Palamas', 'Gregory Palamas'),

('thomas-aquinas', 'Thomas Aquinas', 'Friar, Doctor of the Church', 1225, 1274, 'medieval', 'catholic',
 'The "Angelic Doctor." His Summa Theologiae is one of the most influential works in Western theology and philosophy.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Thomas_Aquinas', 'Thomas Aquinas'),

('francis-of-assisi', 'Francis of Assisi', 'Founder, Order of Friars Minor', 1182, 1226, 'medieval', 'catholic',
 'Beloved saint known for radical poverty, love of creation, and bearing the stigmata. Founded the Franciscan order.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Francis_of_Assisi', 'Francis of Assisi'),

-- =============================================
-- TIER 5: REFORMATION & EARLY MODERN
-- =============================================

('theophanes-the-recluse', 'Theophan the Recluse', 'Bishop', 1815, 1894, 'modern', 'orthodox',
 'Russian bishop who withdrew to monastic seclusion. Translated the Philokalia into Russian. Prolific spiritual writer.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Theophan_the_Recluse', 'Theophan the Recluse'),

('seraphim-of-sarov', 'Seraphim of Sarov', 'Monk, Wonderworker', 1754, 1833, 'modern', 'orthodox',
 'One of the most venerated Russian saints. Greeted all visitors with "My joy, Christ is risen!" Taught that the goal of Christian life is the acquisition of the Holy Spirit.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Seraphim_of_Sarov', 'Seraphim of Sarov'),

-- =============================================
-- TIER 6: MODERN ORTHODOX SAINTS & ELDERS
-- =============================================

('paisios-of-mount-athos', 'Paisios of Mount Athos', 'Monk, Elder', 1924, 1994, 'contemporary', 'orthodox',
 'Beloved Athonite elder known for his humor, wisdom, and gift of discernment. Canonized in 2015. His spiritual counsels have been published in multiple volumes.',
 true, 'short_quote_fair_use', 'https://en.wikipedia.org/wiki/Paisios_of_Mount_Athos', 'Paisios of Mount Athos'),

('porphyrios-of-kavsokalyvia', 'Porphyrios of Kavsokalyvia', 'Hieromonk, Elder', 1906, 1991, 'contemporary', 'orthodox',
 'Athonite elder with the gift of clairvoyance from age 12. Known for his joyful approach to spiritual life. Canonized in 2013.',
 true, 'short_quote_fair_use', 'https://en.wikipedia.org/wiki/Porphyrios_of_Kavsokalyvia', NULL),

('nektarios-of-aegina', 'Nektarios of Aegina', 'Metropolitan of Pentapolis', 1846, 1920, 'modern', 'orthodox',
 'Miracle worker and author of theological and pastoral works. Founded the Holy Trinity Convent on Aegina. Canonized in 1961.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Nectarius_of_Aegina', 'Nectarius of Aegina'),

('cleopa-of-sihastria', 'Cleopa Ilie', 'Archimandrite, Elder', 1912, 1998, 'contemporary', 'orthodox',
 'Romanian elder of Sihăstria Monastery. Renowned spiritual father who shepherded thousands. His teachings are treasured in Romanian Orthodoxy.',
 false, 'short_quote_fair_use', 'https://en.wikipedia.org/wiki/Cleopa_Ilie', NULL),

('sophrony-of-essex', 'Sophrony Sakharov', 'Archimandrite', 1896, 1993, 'contemporary', 'orthodox',
 'Disciple of St. Silouan the Athonite. Founded the Patriarchal Stavropegic Monastery of St. John the Baptist in Essex, England. Author of "His Life Is Mine."',
 true, 'short_quote_fair_use', 'https://en.wikipedia.org/wiki/Sophrony_(Sakharov)', NULL),

('silouan-the-athonite', 'Silouan the Athonite', 'Monk', 1866, 1938, 'contemporary', 'orthodox',
 'Russian monk of the St. Panteleimon Monastery on Mount Athos. Known for his profound humility and teaching on keeping the mind in hell without despair.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Silouan_the_Athonite', 'Silouan the Athonite'),

('john-maximovitch', 'John of Shanghai and San Francisco', 'Archbishop', 1896, 1966, 'contemporary', 'orthodox',
 'Wonderworker of the modern age. Served in Shanghai during WWII, saving refugees. Known for extreme asceticism and walking barefoot.',
 true, 'short_quote_fair_use', 'https://en.wikipedia.org/wiki/John_of_Shanghai_and_San_Francisco', 'John of Shanghai and San Francisco'),

('nikolaj-velimirovic', 'Nikolaj Velimirović', 'Bishop of Ohrid and Žiča', 1881, 1956, 'contemporary', 'orthodox',
 'Serbian bishop, theologian, and poet. Survived Dachau concentration camp. Author of "Prayers by the Lake" and the Prologue of Ohrid.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Nikolaj_Velimirovi%C4%87', NULL),

('justin-popovic', 'Justin Popović', 'Archimandrite', 1894, 1979, 'contemporary', 'orthodox',
 'Serbian theologian called "the new Chrysostom." Author of a comprehensive dogmatics. Canonized in 2010.',
 true, 'short_quote_fair_use', 'https://en.wikipedia.org/wiki/Justin_Popovi%C4%87', NULL),

('thaddeus-of-vitovnica', 'Thaddeus of Vitovnica', 'Archimandrite', 1914, 2003, 'contemporary', 'orthodox',
 'Serbian elder known for "Our Thoughts Determine Our Lives." His simple, profound teachings on inner peace have reached millions.',
 false, 'short_quote_fair_use', 'https://en.wikipedia.org/wiki/Thaddeus_of_Vitovnica', NULL),

('joseph-the-hesychast', 'Joseph the Hesychast', 'Monk, Elder', 1897, 1959, 'contemporary', 'orthodox',
 'Athonite cave-dwelling elder who revived hesychastic practice on Mount Athos. His disciples became abbots of major Athonite monasteries.',
 true, 'short_quote_fair_use', 'https://en.wikipedia.org/wiki/Joseph_the_Hesychast', NULL),

('philaret-of-moscow', 'Philaret of Moscow', 'Metropolitan', 1782, 1867, 'modern', 'orthodox',
 'Leading figure of Russian theology in the 19th century. Oversaw the Russian Bible translation. Known for theological precision.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/Philaret_(Drozdov)', NULL),

('john-of-kronstadt', 'John of Kronstadt', 'Priest', 1829, 1908, 'modern', 'orthodox',
 'Parish priest renowned for his fervent liturgical celebrations and charitable work. His diary "My Life in Christ" is a spiritual classic.',
 true, 'public_domain', 'https://en.wikipedia.org/wiki/John_of_Kronstadt', 'John of Kronstadt')
ON CONFLICT (slug) DO NOTHING;
