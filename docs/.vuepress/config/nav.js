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
          { text: '多线程', link: '/java/se/thread/study-note/' },

        ]
      },
      {
        text: 'Java-ee', items: [
          { text: 'MySQL', link: '/mysql/overview/' },
          { text: 'JavaWeb', link: '/javaweb/basic-concepts/' },
        ]
      },

    ]
  },
  {
    text: 'Web',
    items: [
      { text: 'HTML', link: '/html5/overview/' },
      { text: 'CSS', link: '/css/overview/' },
      { text: 'JavaScript', link: '/javascript/overview/' },
      { text: 'vue', link: '/vue/overview/' },
    ]
  },
  {
    text: '框架',
    items: [
      { text: 'MyBatis', link: '/mybatis/study-note/' },
      { text: 'MyBatis-Plus', link: '/mybatis-plus/study-note/' },
	  { text: 'Spring', link: '/spring/study-note/' },
      { text: 'Redis', link: '/redis/study-note/' },
	  
	  
    ]
  },
  {
    text: '其他',
    items: [
      { text: 'C语言', link: '/c/' },
	  { text: 'git', link: '/git/' },
{ text: 'SMC&P', link: '/SMC&P/note/' },
    ]
  },
  { text: '关于', link: '/about/' },
  { text: '归档', link: '/archives/' }
]