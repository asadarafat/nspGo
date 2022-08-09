(function (nx) {

    // data = JSON.parse(data) 
    // data loaded from index.html via nextData.js

    console.log(data)
    var activeLayout = ''
    var defaultIconType = 'switch'

    var topo = new nx.graphic.Topology({
        showIcon: true,
        adaptive: true,
        nodeConfig: {
            // label: 'model.name',
            label: 'model.id',
            iconType: function (model) {
                if (model._data.group === 'N/A') {
                    return defaultIconType
                } else {
                    return model._data.group
                }
            },
        },
        linkConfig: {
            linkType: 'parallel',
            width: 2,
            sourcelabel: 'model.source_endpoint',
            targetlabel: 'model.target_endpoint',
            color: '#DBEAFE',
        },
        identityKey: 'id',
        // dataProcessor: 'force',
        // sort order consists of typical Clos hierarchy levels mixed with numerical values to help achieve auto sorting on arbitrary topologies
        layoutConfig: {
            sortOrder: ['10', '9', 'superspine', '8', 'dc-gw', '7', '6', 'spine', '5', '4', 'leaf', 'border-leaf', '3', 'server', '2', '1'],
        },
        enableSmartLabel: true,
        enableSmartNode: true,
        enableGradualScaling: true,
        autoLayout: true,
        linkInstanceClass: 'CustomLinkLabel',
        supportMultipleLink: true,
        tooltipManagerConfig: {
            nodeTooltipContentClass: 'CustomNodeTooltip',
        },

    });

    // when group property is not set in containerlab
    // we set it to N/A to indicate a missing group assignment
    for (var key in data.nodes) {
        if (!("group" in data.nodes[key])) {
            data.nodes[key]["group"] = 'N/A';
        }
    }

    nx.define('CustomLinkLabel', nx.graphic.Topology.Link, {
        properties: {
            sourcelabel: 'null',
            targetlabel: 'null',
        },
        view: function (view) {
            view.content.push({
                    name: 'sourceBadge',
                    type: 'nx.graphic.Group',
                    content: [{
                            name: 'sourceBg',
                            type: 'nx.graphic.Rect',
                            props: {
                                'class': 'link-set-circle',
                                height: 1
                            }
                        },
                        {
                            name: 'sourceText',
                            type: 'nx.graphic.Text',
                            props: {
                                'class': 'link-set-text',
                                y: 1
                            }
                        }
                    ],
                    props: {
                        'alignment-baseline': 'after-edge',
                    }
                }, {
                    name: 'targetBadge',
                    type: 'nx.graphic.Group',
                    content: [{
                            name: 'targetBg',
                            type: 'nx.graphic.Rect',
                            props: {
                                'class': 'link-set-circle',
                                height: 1
                            }
                        },
                        {
                            name: 'targetText',
                            type: 'nx.graphic.Text',
                            props: {
                                'class': 'link-set-text',
                                y: 1
                            }
                        }
                    ],
                    props: {
                        'alignment-baseline': 'after-edge',
                    }
                }

            );
            return view;
        },
        methods: {
            init: function (args) {
                this.inherited(args);
                this.topology().fit();
            },
            'setModel': function (model) {
                this.inherited(model);
            },
            update: function () {
                this.inherited();
                var line = this.line();
                var angle = line.angle();
                var stageScale = this.stageScale();
                line = line.pad(50 * stageScale, 50 * stageScale);
                if (this.sourcelabel()) {
                    var sourceBadge = this.view('sourceBadge');
                    var sourceText = this.view('sourceText');
                    var sourceBg = this.view('sourceBg');
                    var point;
                    sourceText.sets({
                        text: this.sourcelabel(),
                    });
                    //TODO: accommodate larger text label
                    sourceBg.sets({
                        width: 50,
                        visible: true
                    });
                    sourceBg.setTransform(50 / -2);
                    point = line.start;
                    if (stageScale) {
                        sourceBadge.set('transform', 'translate(' + point.x + ',' + point.y + ') ' + 'scale (' + stageScale + ') ');
                    } else {
                        sourceBadge.set('transform', 'translate(' + point.x + ',' + point.y + ') ');
                    }
                }
                if (this.targetlabel()) {
                    var targetBadge = this.view('targetBadge');
                    var targetText = this.view('targetText');
                    var targetBg = this.view('targetBg');
                    var point;
                    targetText.sets({
                        text: this.targetlabel(),
                    });
                    targetBg.sets({
                        width: 50,
                        visible: true
                    });
                    targetBg.setTransform(50 / -2);
                    point = line.end;
                    if (stageScale) {
                        targetBadge.set('transform', 'translate(' + point.x + ',' + point.y + ') ' + 'scale (' + stageScale + ') ');
                    } else {
                        targetBadge.set('transform', 'translate(' + point.x + ',' + point.y + ') ');
                    }
                }
                this.view("sourceBadge").visible(false);
                this.view("sourceBg").visible(false);
                this.view("sourceText").visible(false);
                this.view("targetBadge").visible(false);
                this.view("targetBg").visible(false);
                this.view("targetText").visible(false);
            }
        }
    });

    nx.define('CustomNodeTooltip', nx.ui.Component, {
        properties: {
            node: {},
            topology: {}
        },
        view: {
            tag: 'div',
            content: [{
                    tag: 'div',
                    content: '{#node.model.name}',
                    props: {
                        "class": "font-bold text-black text-center uppercase border-b pb-2"
                    },
                },
                {
                    tag: 'div',
                    content: [{
                            tag: 'div',
                            content: [{
                                    tag: 'label',
                                    content: 'Image: ',
                                    props: {
                                        "class": "font-semibold text-black pt-2"
                                    },
                                },
                                {
                                    tag: 'label',
                                    content: 'Kind: ',
                                    props: {
                                        "class": "font-semibold text-black"
                                    },
                                },
                                {
                                    tag: 'label',
                                    content: 'Group: ',
                                    props: {
                                        "class": "font-semibold text-black"
                                    },
                                },
                                {
                                    tag: 'label',
                                    content: 'State: ',
                                    props: {
                                        "class": "font-semibold text-black"
                                    },
                                },
                                {
                                    tag: 'label',
                                    content: 'IPv4: ',
                                    props: {
                                        "class": "font-semibold text-black"
                                    },
                                },
                                {
                                    tag: 'label',
                                    content: 'IPv6: ',
                                    props: {
                                        "class": "font-semibold text-black"
                                    },
                                },
                            ],
                            props: {
                                "class": "flex flex-col pr-3"
                            },
                        },
                        {
                            tag: 'div',
                            content: [{
                                    tag: 'span',
                                    content: '{#node.model.image}',
                                    props: {
                                        "class": "font-normal text-black pt-2 inline-table"
                                    },
                                },
                                {
                                    tag: 'span',
                                    content: '{#node.model.kind}',
                                    props: {
                                        "class": "font-normal text-black"
                                    },
                                },
                                {
                                    tag: 'span',
                                    content: '{#node.model.group}',
                                    props: {
                                        "class": "font-normal text-black"
                                    },
                                },
                                {
                                    tag: 'span',
                                    content: '{#node.model.state}',
                                    props: {
                                        "class": "font-normal text-black"
                                    },
                                },
                                {
                                    tag: 'span',
                                    content: '{#node.model.ipv4_address}',
                                    props: {
                                        "class": "font-normal text-black"
                                    },
                                },
                                {
                                    tag: 'span',
                                    content: '{#node.model.ipv6_address}',
                                    props: {
                                        "class": "font-normal text-black"
                                    },
                                },
                            ],
                            props: {
                                "class": "flex flex-col"
                            },
                        },
                    ],
                    props: {
                        "class": "inline-flex"
                    },
                },
            ],
            props: {
                "class": "bg-white text-sm whitespace-nowrap"
            }
        },
    });

    adaptToContainer = function () {
        topo.adaptToContainer();
    };

    horizontal = function () {
        // document.getElementById("v-btn").classList.remove("text-white", "bg-[#0386d2]");
        // document.getElementById("a-btn").classList.remove("text-white", "bg-[#0386d2]");
        // document.getElementById("h-btn").classList.add("text-white", "bg-[#0386d2]");

        var groupsLayer = topo.getLayer('groups')._content._data;

        console.log(groupsLayer)
        var links = topo.getLayer("links").links()
        console.log(links)


        // if (activeLayout === 'horizontal') {
        //     return;
        // };

        // document.getElementById("a-btn").classList.add("text-white", "bg-[#0386d2]");
        activeLayout = 'horizontal';
        var layout = topo.getLayout('hierarchicalLayout');
        layout.direction('horizontal');
        layout.levelBy(function (node, model) {
            return model.get('group');
        });
        topo.activateLayout('hierarchicalLayout');

        console.log(layout)
    }

    // vertical = function () {
    //     document.getElementById("h-btn").classList.remove("text-white", "bg-[#0386d2]");
    //     document.getElementById("a-btn").classList.remove("text-white", "bg-[#0386d2]");
    //     document.getElementById("v-btn").classList.add("text-white", "bg-[#0386d2]");
    //     if (activeLayout === 'vertical') {
    //         return;
    //     };

    //     // document.getElementById("a-btn").classList.add("text-white", "bg-[#0386d2]");

    //     activeLayout = 'vertical';
    //     var layout = topo.getLayout('hierarchicalLayout');
    //     layout.direction('vertical');
    //     layout.levelBy(function (node, model) {
    //         return model.get('group');
    //     });
    //     topo.activateLayout('hierarchicalLayout');

    // }

    autoLayout = function () {
        var groupsLayer = topo.getLayer('groups')._content._data;

        console.log(groupsLayer)
        var links = topo.getLayer("links").links()
        console.log(links)

        activeLayout = 'force';
        var layout = topo.getLayout('force');
        layout.direction('horizontal');
        layout.levelBy(function (node, model) {
            return model.get('group');
        });
        topo.activateLayout('force');

    }

    asad = function () {
        // document.getElementById("h-btn").classList.remove("text-white", "bg-[#0386d2]");
        // document.getElementById("v-btn").classList.remove("text-white", "bg-[#0386d2]");
        // document.getElementById("a-btn").classList.add("text-white", "bg-[#0386d2]");

        //check if group already added
        if (topo.getLayer('groups')._content._data.length > 0) {
            return;
        }

        //get interTopoLayerLink

        var links = topo.getLayer("links").links()
        // console.log(links[12]._model._data.type)
        // console.log(links.length)
        for (index = 0; index < links.length; index++) {
            console.log(links[index]._model._data.type);
            if (links[index]._model._data.group == "interTopoLayerLink") {
                // console.log(links[index])
                links[index].enable(false)
                // links[index].visible(true)
            }
        }

        var groupsLayer = topo.getLayer('groups');

        // var nodesPhysicalLayer1 = topo.getNode(0);
        // console.log(nodesPhysicalLayer1)

        var nodesAll = topo.getLayer("nodes").nodes()
        // console.log(nodesPhysicalLayer)

        var nodesPhysicalLayer = []
        var nodesIgpLayer = []

        for (index = 0; index < nodesAll.length; index++) {
            if (nodesAll[index]._model._data.group == "PhysicalLayer") {
                // console.log(nodesAll[index])
                nodesPhysicalLayer.push(nodesAll[index])
            }
            if (nodesAll[index]._model._data.group == "IgpLayer") {
                // console.log(nodesAll[index])
                nodesIgpLayer.push(nodesAll[index])
            }
        }

        var groupPhysicalLayer = groupsLayer.addGroup({
            nodes: nodesPhysicalLayer,
            label: 'Layer-2 Topology',
            // color: '#ddffff'
            // color: #ddffdc
            color: '#ddffdc'
        });
        var groupIgpLayer = groupsLayer.addGroup({
            nodes: nodesIgpLayer,
            label: 'Layer-3 Topology-asad',
            color: '#ddffdc'
        });

        var link_layer_el = document.querySelectorAll('[data-nx-type="nx.graphic.Topology.LinksLayer"]')
        var group_layer_el = document.querySelectorAll('[data-nx-type="nx.graphic.Topology.GroupsLayer"]')
        var parent_el = link_layer_el[0].parentNode
        parent_el.insertBefore(group_layer_el[0], link_layer_el[0])
        // topo.insertLayerAfter('mylayer', 'MyLayer', 'nodes');
        // console.log(topo)
    }

    pathDraw = function () {
        var pathLayer = topo.getLayer("paths");
        var links1 = [topo.getLink("from-L3-10.10.10.3-to-L3-10.10.10.5"), topo.getLink("from-L3-10.10.10.5-to-L3-10.10.10.4")];
        console.log(links1)
        var path1 = new nx.graphic.Topology.Path({
            links: links1,
            arrow: 'cap',
            pathGutter: 2,
            pathWidth: 5,
        });
        pathLayer.addPath(path1);
        var links2 = [topo.getLink("from-L3-10.10.10.1-to-L3-10.10.10.2")];
        console.log(links2)
        var path2 = new nx.graphic.Topology.Path({
            links: links2,
            arrow: 'end',
            pathGutter: 2,
            pathWidth: 5,
        });
        pathLayer.addPath(path2);
    }

    window.onresize = adaptToContainer;
    var app = new nx.ui.Application();
    app.container(document.getElementById('clab-topology'));
    topo.attach(app);
    topo.on('ready', function () {
        topo.data(data);
        asad();
        pathDraw();
    });


})(nx);