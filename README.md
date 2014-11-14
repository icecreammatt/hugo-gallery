# [Hugo](http://hugo.spf13.com) image gallery generator

This tool will create a new posts directory containing a markdown file for each image in source directory allowing for an ordered slide show.

## Usage
`hugo-gallery <Source Path> <Destination Section> <Title>`

## Example

`hugo-gallery static/images/vacation-photos hawaii "Hawaii Trip"`

Visit `localhost:1313/hawaii` to view the content.

This would read all of the images out of the `static/images/vacation-photos` directory and create a new folder named `hawaii` in `content/hawaii` filled with front matter markdown files. See sample below for details.

### Markdown Sample

```yml
---
title: Hawaii Trip
date: "2014-11-12"
image_name: images/vacation-photos/IMG_003.jpg
previous_image: images/vacation-photos/IMAGE_002.jpg
next_image: images/vacation-photos/IMAGE_004.jpg
next_post_path: hawaii/IMAGE_004
previous_post_path: hawaii/IMAGE_002
---
```

### Template Usage
Reference `image_name` in the Hugo single post template  
`<img src="{{ .Params.image_name }}" />`

## Todo:
* Implement test coverage to solution

## License
* MIT
