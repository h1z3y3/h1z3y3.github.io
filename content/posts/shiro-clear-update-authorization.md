---
title: Shiro清除更新缓存的用户权限
date: 2017-04-04 18:57:44
tags: [J2EE, Shiro]
---

Apache Shiro用于权限管理十分方便，但存在一个问题，就是当用户的权限发生变化的时候，就需要用户重新登录，重新缓存用户的权限信息。  
现在想要在改变用户的权限的时候，清理用户的权限。    
在写的过程中查找了一些资料，但是并没有成功实现权限的清理，所以我进行了一些修改，并实现了Helper类。

<!--more-->

我写的帮助类：

	
	package com.zhaoyang.core.feature.security;
	
	import org.apache.shiro.SecurityUtils;
	import org.apache.shiro.cache.Cache;
	import org.apache.shiro.cache.CacheManager;
	import org.apache.shiro.subject.SimplePrincipalCollection;
	import org.apache.shiro.subject.Subject;
	import org.slf4j.Logger;
	import org.slf4j.LoggerFactory;
	
	/**
	 * @author zhaoyang
	 * 
	 * @since 2017-04-04 2:12 PM
	 */
	public class ShiroAuthorizationHelper {
	
	    //shiro 配置的cacheManager， 需要使用Spring bean进行注入
	    private static CacheManager cacheManager;
	
	    private static Logger logger = LoggerFactory.getLogger(ShiroAuthorizationHelper.class);
	
	    /**
	     * 清除用户的权限
	     * 
	     * 这里需要注意的是，
	     * 网上很多实现都是这里只传递一个String类型的username过来，根据这个String当key去清除Cache
	     * 但是Shiro在缓存用户权限的时候使用的key并不是String类型，所以调用remove的时候并不能清除缓存的权限
	     *
	     * shiro缓存时使用的key，是登录时使用的SimplePrincipalCollection对象，所以remove的时候需要的不是一个String值，
	     * 具体可以参考下面方法中打印cache的key的过程, 可以看到打印出key的类是 `class org.apache.shiro.subject.SimplePrincipalCollection`
	     * 所以你cache.remove(String username)肯定清除不了
	     *
	     * @param principal
	     */
	    public static void clearAuthorizationInfo(SimplePrincipalCollection principal) {
	        logger.info("clear the user: " + principal.toString() + "'s authorizationInfo");
	        Cache<Object, Object> cache = cacheManager.getCache("myShiroCache");//myShiroCache是我配置用于缓存的cache的Name，在spring配置文件中配置，可以看文章最后
	
	//        for (Object k : cache.keys()) {
	//            System.out.println(k.getClass());
	//        }
	
	        cache.remove(principal);
	
	    }
	
	    /**
	     * 清除当前用户的权限
	     */
	    public static void clearAuthorizationInfo() {
	        if (SecurityUtils.getSubject().isAuthenticated()) {
	            Subject subject = SecurityUtils.getSubject();
	            String username = subject.getPrincipal().toString();
	            String realmName = subject.getPrincipals().getRealmNames().iterator().next();
	            SimplePrincipalCollection principalCollection = new SimplePrincipalCollection(username, realmName);
	            // 调用清理用户权限
	            clearAuthorizationInfo(principalCollection);
	        }
	    }
	
	    /**
	     * 由Spring bean将对象注入
	     * @param cacheManager
	     */
	    public static void setCacheManager(CacheManager cacheManager) {
	        ShiroAuthorizationHelper.cacheManager = cacheManager;
	    }
	
	
	}


将cacheManager注入到帮助类：   

	 <!-- 注入Shiro帮助类的cacheManager -->
	    <bean class="org.springframework.beans.factory.config.MethodInvokingFactoryBean">
	        <property name="staticMethod" value="com.damaiya.DMYSite.core.feature.security.ShiroAuthorizationHelper.setCacheManager"/>
	        <property name="arguments" ref="shiroEhcacheManager"/>
	    </bean>
    
      
  
当然ref=“shiroEhcacheManager”需要你自己去实现, 我这里贴下我的：  


	 <!-- 缓存管理器 使用Ehcache实现 -->
	    <bean id="shiroEhcacheManager" class="org.apache.shiro.cache.ehcache.EhCacheManager">
	        <property name="cacheManagerConfigFile" value="classpath:ehcache-shiro.xml"/>
	    </bean>

  

下面是ehcache-shiro.xml配置, 具体的参数作用我就不说了:
  
	<ehcache updateCheck="false" name="shiroCache">
	    <defaultCache
	            maxElementsInMemory="10000"
	            eternal="false"
	            timeToIdleSeconds="120"
	            timeToLiveSeconds="120"
	            overflowToDisk="false"
	            diskPersistent="false"
	            diskExpiryThreadIntervalSeconds="120"
	    />
	
	    <cache name="myShiroCache"
	           maxElementsInMemory="10000"
	           eternal="false"
	           timeToIdleSeconds="30"
	           timeToLiveSeconds="0"
	           overflowToDisk="false"
	           diskPersistent="false"
	           diskExpiryThreadIntervalSeconds="120"/>
	</ehcache>

  

使用指定Name的Cache进行权限缓存配置, securityRealm是我自己的Realm：    


	<bean id="securityRealm" class="com.damaiya.DMYSite.web.security.SecurityRealm"> 
	        <property name="authorizationCacheName" value="myShiroCache"/>    
	</bean>


  
这样整个配置就完成了，而且调用`clearAuthorizationInfo()`时就可以清除当前登录用户的权限信息了。
