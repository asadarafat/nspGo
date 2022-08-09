var topologyData = {
	"nodes": [
		//Physical
		{
			"id": 11,
			"name": "10.0.0.1",
			"layer": "Physical"
		},
		{
			"id": 12,
			"name": "10.0.0.2",
			"layer": "Physical"
		},
		{
			"id": 13,
			"name": "10.0.0.3",
			"layer": "Physical"
		},
		{
			"id": 14,
			"name": "10.0.0.4",
			"layer": "Physical"
		},
		{
			"id": 15,
			"name": "10.0.0.5",
			"layer": "Physical"
		},
		//IGP
		{
			"id": 21,
			"name": "10.0.0.1",
			"layer": "IGP"
		},
		{
			"id": 22,
			"name": "10.0.0.2",
			"layer": "IGP"
		},
		{
			"id": 23,
			"name": "10.0.0.3",
			"layer": "IGP"
		},
		{
			"id": 24,
			"name": "10.0.0.4",
			"layer": "IGP"
		},
		{
			"id": 25,
			"name": "10.0.0.5",
			"layer": "IGP"
		},
		//MPLS
    {
      "id": 31,
      "name": "10.0.0.1",
      "layer": "MPLS"
    },
    {
      "id": 32,
      "name": "10.0.0.2",
      "layer": "MPLS"
    },
    {
      "id": 33,
      "name": "10.0.0.3",
      "layer": "MPLS"
    },
    {
      "id": 34,
      "name": "10.0.0.4",
      "layer": "MPLS"
    },
    {
      "id": 35,
      "name": "10.0.0.5",
      "layer": "MPLS"
    },
		//Service
    {
      "id": 41,
      "name": "10.0.0.1",
      "layer": "Service"
    },
    {
      "id": 42,
      "name": "10.0.0.2",
      "layer": "Service"
    },
    {
      "id": 43,
      "name": "10.0.0.3",
      "layer": "Service"
    },
    {
      "id": 44,
      "name": "10.0.0.4",
      "layer": "Service"
    },
    {
      "id": 45,
      "name": "10.0.0.5",
      "layer": "Service"
    }
	],
	"links": [
		//Physical
		{
			"source": 11,
			"target": 12
		},
		{
			"source": 11,
			"target": 14
		},
		{
			"source": 15,
			"target": 14
		},
		{
			"source": 13,
			"target": 12
		},
		{
			"source": 15,
			"target": 13
		},
		//IGP
		{
			"source": 21,
			"target": 22
		},
		{
			"source": 21,
			"target": 24
		},
		{
			"source": 25,
			"target": 24
		},
		{
			"source": 23,
			"target": 22
		},
		{
			"source": 25,
			"target": 23
		},
		//MPLS
    {
      "source": 31,
      "target": 32
    },
    {
      "source": 31,
      "target": 34
    },
    {
      "source": 35,
      "target": 34
    },
    {
      "source": 33,
      "target": 32
    },
    {
      "source": 35,
      "target": 33
    },
		//Service
    {
      "source": 41,
      "target": 42
    },
    {
      "source": 41,
      "target": 44
    },
    {
      "source": 45,
      "target": 44
    },
    {
      "source": 43,
      "target": 42
    },
    {
      "source": 45,
      "target": 43
    },
		//Physical-IGP
		{
			"source": 11,
			"target": 21
		},
		{
			"source": 12,
			"target": 22
		},
		{
			"source": 13,
			"target": 23
		},
		{
			"source": 14,
			"target": 24
		},
		{
			"source": 15,
			"target": 25
		},
		//IGP-MPLS
    {
      "source": 21,
      "target": 31
    },
    {
      "source": 22,
      "target": 32
    },
    {
      "source": 23,
      "target": 33
    },
    {
      "source": 24,
      "target": 34
    },
    {
      "source": 25,
      "target": 35
    },
		//MPLS-Service
    {
      "source": 31,
      "target": 41
    },
    {
      "source": 32,
      "target": 42
    },
    {
      "source": 33,
      "target": 43
    },
    {
      "source": 34,
      "target": 44
    },
    {
      "source": 35,
      "target": 45
    }
	]
};