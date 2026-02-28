-- Seed data: Topics
-- These are the core theological topics for categorizing patristic and saintly quotes.

INSERT INTO topics (slug, name, description) VALUES
  ('trinity', 'The Holy Trinity', 'Quotes on the nature of the Triune God: Father, Son, and Holy Spirit'),
  ('christology', 'Christology', 'The person, nature, and work of Jesus Christ'),
  ('salvation', 'Salvation & Theosis', 'Soteriology: redemption, grace, and the path to deification'),
  ('prayer', 'Prayer & Contemplation', 'On prayer, meditation, and the inner life'),
  ('repentance', 'Repentance', 'Metanoia, confession, and the turning of the heart toward God'),
  ('humility', 'Humility', 'The mother of all virtues'),
  ('love', 'Love & Charity', 'Divine and human love, agape, compassion'),
  ('faith', 'Faith', 'Trust in God, the substance of things hoped for'),
  ('suffering', 'Suffering & Patience', 'The meaning of suffering and endurance in Christ'),
  ('eucharist', 'The Eucharist', 'The Body and Blood of Christ, the mystical supper'),
  ('scripture', 'Holy Scripture', 'On reading, interpreting, and living the Word of God'),
  ('church', 'The Church', 'Ecclesiology: the Body of Christ, unity, and tradition'),
  ('monasticism', 'Monastic Life', 'The ascetic path, desert wisdom, and spiritual discipline'),
  ('creation', 'Creation & Nature', 'God as Creator, stewardship, and the natural world'),
  ('death', 'Death & Resurrection', 'Eschatology, the afterlife, and the hope of resurrection'),
  ('virtue', 'Virtue & Holiness', 'The pursuit of virtue and the life of holiness'),
  ('sin', 'Sin & Temptation', 'The nature of sin and the struggle against temptation'),
  ('theotokos', 'The Theotokos', 'The Ever-Virgin Mary, Mother of God'),
  ('saints', 'The Saints', 'The communion of saints, intercession, and holy examples'),
  ('wisdom', 'Wisdom & Knowledge', 'Divine wisdom, spiritual discernment, and the fear of the Lord'),
  ('peace', 'Peace & Stillness', 'Hesychia, inner peace, and serenity in God'),
  ('obedience', 'Obedience', 'Submission to God''s will and spiritual authority'),
  ('fasting', 'Fasting & Asceticism', 'Bodily discipline as a path to spiritual growth'),
  ('joy', 'Joy', 'Spiritual joy and rejoicing in the Lord'),
  ('icons', 'Icons & Sacred Art', 'The theology and veneration of holy images')
ON CONFLICT (slug) DO NOTHING;
