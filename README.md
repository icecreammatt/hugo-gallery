# [Hugo](http://hugo.spf13.com) image gallery generator

This tool will create a new posts directory containing a markdown file for each image in source directory allowing for an ordered slide show.

## Usage
`hugo-gallery <Source Path> <Destination Section> <Title>`

## Example

`hugo-gallery static/images/vacation-photos hawaii "Hawaii Trip"`

This would read all of the images out of the `static/images/vacation-photos` directory and create a new folder named `hawaii` in `content/hawaii` filled with front matter markdown files which contain a `title` of `Hawaii Trip` a `weight` matching the images index and a `image_name` variable for use in a template.

Visit `localhost:1313/hawaii` to view the content.

### Markdown Sample

```yml
---
title: Hawaii Trip
date: "2014-11-12"
weight: 3
image_name: images/vacation-photos/IMG_003.jpg
---
```

### Template Usage
Reference `image_name` in the Hugo single post template  
`<img src="{{ .Params.image_name }}" />`

## Todo:
* Implement test coverage to solution

## License
* MIT