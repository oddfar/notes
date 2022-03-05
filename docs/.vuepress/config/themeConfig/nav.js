module.exports = [
  { text: '首页', link: '/' },
  {
    text: 'Java', link: '/java/', items: [
      {
        text: 'Java-Se', items: [
          { text: 'Java基础', link: '/java/se/initial-java/' },
          { text: 'Java面向对象', link: '/java/se/object/' },
          { text: 'Java常用类', link: '/java/se/commonly-used-class/' },
          { text: 'Java集合框架', link: '/java/se/collection/synopsis/' },

        ]
      },
      {
        text: 'Java-Se进阶', items: [
          { text: 'JUC多线程', link: '/java/se/thread/1/' },

        ]
      },
      {
        text: 'Java-ee', items: [

          { text: 'JavaWeb', link: '/javaweb/' },
        ]
      },

    ]
  },
  {
    text: '数据库',
    items: [

      {
        text: 'SQL 数据库', items: [
          { text: 'MySQL', link: '/mysql/' },

        ]
      },

      {
        text: 'NoSQL 数据库', items: [
          { text: 'Redis', link: '/redis/' },
          { text: 'ElasticSearch', link: '/es/' },
          { text: 'MongoDB', link: '/mongodb/' },
        ]
      },
    ]
  },
  {
    text: '框架',
    items: [
      {
        text: '数据库', items: [
          { text: 'MyBatis', link: '/mybatis/study-note/' },
          { text: 'MyBatis-Plus', link: '/mybatis-plus/study-note/' },
        ]


      }, {
        text: 'Spring', items: [
          { text: 'Spring', link: '/spring/' },
        ]
      }
      , {
        text: '中间件', items: [
          { text: 'RabbitMQ', link: '/rabbitmq/' },
        ]
      }
    ]
  },
  {
    text: '工具|部署',
    items: [
      { text: 'Git', link: '/git/' },
      { text: 'Docker', link: '/docker/' },
      { text: 'Jenkins', link: '/jenkins/' },
    ]
  },
    {
    text: '项目',
    items: [
      { text: '校园信息墙', link: '/campus/' },
    ]
  },
  { text: '关于', link: '/about/' },
  { text: '归档', link: '/archives/' }
]