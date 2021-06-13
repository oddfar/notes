module.exports = [
  { text: '首页', link: '/' },
  {
    text: 'Java', link: '/java/', items: [
      {
        text: 'Java-Se', items: [
          { text: '前言', link: '/java/se/preface/' },
          { text: 'Java基础', link: '/java/se/initial-java/' },
          { text: 'Java面向对象', link: '/java/se/object/' },
          { text: 'Java常用类', link: '/java/se/commonly-used-class/' },
          { text: 'Java集合框架', link: '/java/se/aggregate/synopsis/' },

        ]
      },
      {
        text: 'Java-Se进阶', items: [
          { text: 'JUC多线程', link: '/java/se/thread/study-note/' },

        ]
      },
      {
        text: 'Java-ee', items: [

          { text: 'JavaWeb', link: '/javaweb/basic-concepts/' },
        ]
      },

    ]
  },
  {
    text: '数据库',
    items: [

      {
        text: 'SQL 数据库', items: [
          { text: 'MySQL', link: '/mysql/preface/' },

        ]
      },

      {
        text: 'NoSQL 数据库', items: [
          { text: 'Redis', link: '/redis/study-note/1/' },
          { text: 'ElasticSearch', link: '/es/' },
        ]
      },
    ]
  },
  {
    text: '框架',
    items: [
      { text: 'MyBatis', link: '/mybatis/study-note/' },
      { text: 'MyBatis-Plus', link: '/mybatis-plus/study-note/' },
      { text: 'Spring', link: '/spring/study-note/' },



    ]
  },
  {
    text: '工具|部署',
    items: [
      { text: 'Git', link: '/git/simple-commit/' },
      { text: 'Docker', link: '/docker/overview/' },
    ]
  },
  {
    text: '其他',
    items: [
      { text: 'C语言', link: '/c/' },
      { text: 'SMC&P', link: '/SMC&P/note/' },
    ]
  },
  { text: '关于', link: '/about/' },
  { text: '归档', link: '/archives/' }
]