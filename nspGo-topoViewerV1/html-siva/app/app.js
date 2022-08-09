var app = new nx.ui.Application();
var topologyConfig = {
	"width": 800,
	"height": 800,
	"identityKey": "id",
	"nodeConfig": {
		"label": "model.id",
		"iconType": "router"
	},
	"linkConfig": {
		"linkType": "curve"
	},
	"showIcon": true,
	"dataProcessor": "force"
};

var topology = new nx.graphic.Topology(topologyConfig);
topology.data(topologyData);

function randomColor() {
	return "#000000".replace(/0/g,function(){return (~~(Math.random()*16)).toString(16);});
}

function updateGroups() {
	var layer = document.querySelectorAll('input[name="layer[]"]:checked');
	var layerChecked = [];
	if(layer.length > 0) {
		for(var i = 0; i < layer.length; i++) {
			layerChecked.push(layer[i].value);
		}
	}
	var groupsLayer = topology.getLayer("groups");
	var nodesDict = topology.getLayer("nodes").nodeDictionary();
	var groups = {};
	var topoNodes = topologyData["nodes"];
	for(var i = 0; i < topoNodes.length; i++) {
		var nodeId = topoNodes[i]["id"];
		var groupName = topoNodes[i]["LayerName"];
		if(groups[groupName] == undefined) {
			groups[groupName] = []
		}
		groups[groupName].push(nodesDict.getItem(nodeId));
	}
	const groupKeys = Object.keys(groups);
	for(var i = 0; i < groupKeys.length; i++) {
		var nodes = [];
		var key = groupKeys[i];
		if(layerChecked.includes(key)) {
			if(!groupsLayer.getGroup(key)) {
				groupsLayer.addGroup({
					nodes: groups[key],
					color: '#ddffff',
					label: key,
					id: key
				});
			} else {
				if(!groupsLayer.getGroup(key).visible()) {
					groupsLayer.getGroup(key).show();
				}
			}
		} else {
			groupsLayer.getGroup(key).hide();
		}
	}
}

topology.on("topologyGenerated", function() {
	updateGroups();
});

topology.attach(app);