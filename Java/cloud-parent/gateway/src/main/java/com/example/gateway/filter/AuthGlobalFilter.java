package com.example.gateway.filter;

import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.util.List;

//自定义全局过滤器需要实现GlobalFilter和Ordered接口
@Component
public class AuthGlobalFilter implements GlobalFilter, Ordered {

    //完成判断逻辑
    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        //String token = exchange.getRequest().getQueryParams().getFirst("token");
        List<String> token1 = exchange.getRequest().getHeaders().get("token");

        if (token1 == null) {
            System.out.println("鉴权失败");
            exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
            return exchange.getResponse().setComplete();
        }

        System.out.println(token1.get(0));

        //调用chain.filter继续向下游执行
        return chain.filter(exchange);
    }

    //顺序,数值越小,优先级越高
    @Override
    public int getOrder() {
        return 0;
    }
}
