/**
 * 创建人: bullub
 * 创建时间: 2015-10-09 14:07
 * 用途:
 */
(function (global, doc, $, undefined) {
    "use strict";

    var loop = function () {};

    //导出ViewController类
    global.ViewController = ViewController;

    /**
     * 将普通对象包装成子类的原型，以便Object.create方法使用
     * @param subClassProto {Object} 子类原型对象
     * @param key {String} 键
     * @private
     */
    var _swapSubClassProto = function (subClassProto, key) {
        if (!subClassProto[key] || !subClassProto[key].value) {
            subClassProto[key] = {
                value: subClassProto[key],
                enumerable: true,
                writable: true,
                configurable: true
            };
        }
    };

    /**
     * 将内容渲染到节点中
     * @param $el 节点
     * @param content 内容
     * @param isAppend 是否追加
     * @private
     */
    var _renderElement = function($el, content, isAppend) {

        if (isAppend) {
            $el.append(content);
        } else {
            $el.html(content);
        }
    };

    /**
     * 使得元素失焦
     * @param el 元素
     * @private
     */
    var _blurActiveEl = function(el) {
        try {
            if(el.tagName.toUpperCase() !== 'INPUT') {
                document.activeElement.blur();
            }
        } catch(e){}
    };


    /**
     * ViewController类视图控制器
     * @param rootElement
     */
    function ViewController(rootElement) {
        var self = this,
            $root,
            events = self.events,
            eventHandlers = self.eventHandlers,
            //表达式中的第一个空格
            _start,
            //拿到表达式中的事件名
            _event,
            //拿到表达式中的选择器
            _selector;
        //缓存所有会用到的模板函数
        self.templates = {};

        if (!rootElement) {
            rootElement = doc.body;
        }

        var _proxy = function (handler) {
            return function (event) {
                _blurActiveEl(this);
                handler.call(self, event, $(this));
            };
        };

        var _delegateEvent = function(expression) {
            //表达式中的第一个空格
            _start = expression.indexOf(' ');
            //拿到表达式中的事件名
            _event = expression.substring(0, _start);
            //拿到表达式中的选择器
            _selector = expression.slice(_start - expression.length + 1);

            $root.delegate(_selector, _event, _proxy(eventHandlers[events[expression]]));
        };

        //初始化根节点
        $root = self.$root = $(rootElement);
        //处理声明式事件绑定
        for (var expression in events) {
            _delegateEvent(expression);
        }

    }

    ViewController.prototype = {
        constructor: ViewController,
        /**
         * 发送一个请求
         * @param name
         * @param options
         */
        request: function request(name, options) {
            
        },
        /**
         * 渲染模版数据
         * @param tplId 模版名称
         * @param data 渲染数据
         * @param element 渲染容器
         * @param isAppend 是否追加，默认不追加
         * @returns {String} 渲染后的html内容
         */
        render: function (tplId, data, element, isAppend) {
            var tpl = this.templates[tplId], html;

            if (!tpl) {
                this.templates[tplId] = tpl = _.template($('#' + tplId).html());
            }

            html = tpl(data);

            _renderElement($(element), html, isAppend);

            return html;
        }
    };



    /**
     * 实现继承
     * @param subClassProto 子类的原型及构造函数
     * @returns {Function} 子类的构造函数
     */
    ViewController.extend = function extend(subClassProto) {
        var ParentClass = this;

        if (!subClassProto || typeof subClassProto.constructor !== 'function') {
            throw new Error("sub class has no constructor");
        }

        //拿到子类的构造函数
        var SubClassConstructor = subClassProto.constructor;

        //从子类原型链上删掉构造器字段
        delete subClassProto.constructor;

        //格式化子类的原型链
        for (var key in subClassProto) {
            _swapSubClassProto(subClassProto, key);
        }
        //_super指向父类的原型链，子类不能通过任何方式去修改父类的原型链，父类的原型链在子类中也不可枚举
        subClassProto._super = {
            value: ParentClass.prototype,
            enumerable: false,
            writable: false,
            configurable: false
        };
        //  _Parent主要用于子类直接调用父类构造器，实现参数继承
        // 子类不能通过任何方式去修改父类的构造器，父类的构造器在子类中也不可枚举
        Object.defineProperty(SubClassConstructor, "_Parent", {
            value: function (context, _args) {
                var args = Array.prototype.slice.call(_args);
                ParentClass.apply(context, args);
            },
            enumerable: false,
            writable: false,
            configurable: false
        });
        //实现原型继承
        SubClassConstructor.prototype = Object.create(ParentClass.prototype, subClassProto);
        //子类的extend方法直接使用父类的
        SubClassConstructor.extend = extend;
        SubClassConstructor.prototype.constructor = SubClassConstructor;

        return SubClassConstructor;
    };

}(window, document, jQuery|Zepto));
