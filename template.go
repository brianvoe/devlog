package devlog

var indexHTML = `
<!DOCTYPE html>
<html>
    <head>
        <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
        <script>
            window.onload = function () {
                new Vue({
                    el: '#app',
                    data() {
                        return {
                            port: '<--.Port-->',
                            levels: ['info', 'debug', 'warn', 'error'],
                            data: [],
                            selected: [],
                            startStop: 'start',
                            startStopInterval: null
                        }
                    },
                    mounted() {
                        this.toggleStartStop()
                        this.getData()
                    },
                    methods: {
                        toggleLevel(level) {
                            if (this.levels.includes(level)) {
                                var li = this.levels.indexOf(level);
                                if (li > -1) { this.levels.splice(li, 1); }
                            } else {
                                this.levels.push(level)
                            }

                            this.getData()
                        },

                        toggleSelected(id) {
                            if (this.selected.includes(id)) {
                                var li = this.selected.indexOf(id);
                                if (li > -1) { this.selected.splice(li, 1) }
                            } else {
                                this.selected.push(id)
                            }
                        },

                        toggleStartStop() {
                            if (this.startStop === 'start') {
                                this.startStop = 'stop'
                                this.startStopInterval = setInterval(() => { this.getData() }, 1000);
                            } else {
                                this.startStop = 'start'
                                clearInterval(this.startStopInterval)
                            }
                        },

                        getData() {
                            fetch("http://localhost:8888/getdata", {
                                method: "POST",
                                body: JSON.stringify({
                                    levels: this.levels
                                })
                            })
                            .then((resp) => {return resp.json()})
                            .then((json) => { Vue.set(this, 'data', json) })
                            .catch((err) => {
                                console.log(err)
                            })
                        },

                        syntaxHighlight(json) {
                            if (typeof json == 'string') { json = JSON.parse(json) }
                            if (typeof json != 'string') { json = JSON.stringify(json, null, 2); }
                            json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
                            return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
                                var cls = 'number';
                                if (/^"/.test(match)) {
                                    if (/:$/.test(match)) { cls = 'key'; } else { cls = 'string'; }
                                } else if (/true|false/.test(match)) {
                                    cls = 'boolean';
                                } else if (/null/.test(match)) {
                                    cls = 'null';
                                }
                                return '<span class="' + cls + '">' + match + '</span>';
                            });
                        },

                        toDate(time) {
                            return (new Date(time/1000000)).toLocaleString().replace(', ', '<br />');
                        }
                    }
                })
            }
        </script>
        <style>
            html, body { color: #666666; padding: 0; margin: 0; height: 100%; max-height: 100%; }
            .app {
                display: flex;
                flex-direction: column;
                height: 100%;
                margin: 0 auto;
                min-width: 400px;
                max-width: 800px;
                overflow: hidden;
            }

            .options {
                flex: 0 0 auto;
                display: flex;
                flex-direction: row;
                padding: 8px 0 8px 0;
            }

            .options .option {
                padding: 8px;
                margin: 0 8px 0 0;
                border: solid 1px #cccccc;
                border-radius: 4px;
                cursor: pointer;
            }

            .options .option.selected.info { background-color: #cce5ff; }
            .options .option.selected.debug { background-color: #d4edda; }
            .options .option.selected.warn { background-color: #fff3cd; }
            .options .option.selected.error { background-color: #f8d7da; }
            .options .option.startstop { background-color: #d4edda; margin: 0 0 0 auto; text-transform: capitalize; }

            .data {
                flex: 1 1 auto;
                overflow: auto;
            }

            .item {
                display: flex;
                flex-direction: row;
                padding: 0 0 0 0;
                margin: 0 0 8px 0;
                border-radius: 4px;
                border: solid 1px #cccccc;
                overflow: auto;
            }

            .item .info { background-color: #cce5ff; }
            .item .debug { background-color: #d4edda; }
            .item .warn { background-color: #fff3cd; }
            .item .error { background-color: #f8d7da; }

            .item .level {
                flex: 0 0 30px;
                width: 30px;
                padding: 8px;
                cursor: pointer;
                box-sizing: border-box;
            }
            .item .code {
                flex: 1 1 auto;
                padding: 8px;
            }
            .item .time {
                flex: 0 0 auto;
                text-align: center;
                overflow: hidden;
                font-size: 90%;
                padding: 8px;
            }

            pre { white-space: normal; padding: 0px; margin: 0px; }
            .item.selected pre { white-space: pre-wrap; }
            .key { color: grey; }
            .string { color: green; }
            .number { color: darkorange; }
            .boolean { color: blue; }
            .null { color: magenta; }
        </style>
    </head>
    <body>
        <div id="app" class="app">
            <div class="options">
                <div class="option info" :class="{selected: this.levels.includes('info')}" @click="toggleLevel('info')">Info</div>
                <div class="option debug" :class="{selected: this.levels.includes('debug')}" @click="toggleLevel('debug')">Debug</div>
                <div class="option warn" :class="{selected: this.levels.includes('warn')}" @click="toggleLevel('warn')">Warn</div>
                <div class="option error" :class="{selected: this.levels.includes('error')}" @click="toggleLevel('error')">Error</div>
                <div class="option startstop" @click="toggleStartStop()">{{startStop}}</div>
            </div>
            <div class="data">
                <div :class="['item', selected.includes(item.id) ? 'selected': '']" v-for="item in data" :key="item.id">
                    <div :class="['level', item.level]" @click="toggleSelected(item.id)">
                        <svg
                            x="0px" y="0px"
                            viewBox="0 0 242.133 242.133">
                            <path fill="#666666" style="opacity: .5;" d="M227.133,83.033c8.283,0,15-6.716,15-15V15c0-8.284-6.717-15-15-15H174.1c-8.284,0-15,6.716-15,15
                                s6.716,15,15,15h16.82l-69.854,69.854L51.213,30h16.82c8.284,0,15-6.716,15-15s-6.716-15-15-15H15C6.717,0,0,6.716,0,15v53.033
                                c0,8.284,6.717,15,15,15c8.285,0,15-6.716,15-15v-16.82l69.854,69.854L30,190.92V174.1c0-8.284-6.715-15-15-15
                                c-8.283,0-15,6.716-15,15v53.033c0,8.284,6.717,15,15,15h53.033c8.284,0,15-6.716,15-15c0-8.284-6.716-15-15-15h-16.82
                                l69.854-69.854l69.854,69.854H174.1c-8.284,0-15,6.716-15,15c0,8.284,6.716,15,15,15h53.033c8.283,0,15-6.716,15-15V174.1
                                c0-8.284-6.717-15-15-15c-8.285,0-15,6.716-15,15v16.82l-69.854-69.854l69.854-69.854v16.82
                                C212.133,76.317,218.848,83.033,227.133,83.033z"/>
                        </svg>                            
                    </div>
                    <div class="code"><pre v-html="syntaxHighlight(item.data)"></pre></div>
                    <div class="time" v-html="toDate(item.created_at)"></div>
                </div>
            </div>
        </div>
    </body>
</html>
`
