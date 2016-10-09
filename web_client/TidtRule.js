/**
 * 定义rule命令规则
 */
module.exports = {
    core: [
        {
            name: 'js',
            value: [
                'underscore/underscore',
                'jquery/jQuery',
                'config/default',
                'config/{ENV}',
                'ViewController'
            ]
        }
    ],
    style: [{
        name: "css",
        value: [
            'reset',
            'style'
        ]
    }],
    head: [
        {
            name: 'tpl',
            value: [
                'common/head'
            ]
        },
        {
            name: 'rule',
            value: [
                'style'
            ]
        }
    ]
};