<?php
//Do some stuff before html

    require_once("producer.php");

    queue(["SERVER"=>$_SERVER,"GET"=>$_GET,"POST"=>$_POST,"COOKIE"=>$_COOKIE]);

?>
<!doctype html>

<html lang="en">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>A Basic Sample</title>
    <meta name="description" content="A simple php producer with golang consumer.">
    <meta name="author" content="Luis Rodriguez">

    <link rel="icon" href="/favicon.ico">
    <link rel="icon" href="/favicon.svg" type="image/svg+xml">
    <link rel="apple-touch-icon" href="/apple-touch-icon.png">

    <link rel="stylesheet" href="main.css?v=<?php echo time(); ?>">

</head>

<body>
    Basic Sample

    <script src="https://code.jquery.com/jquery-3.6.0.slim.min.js" integrity="sha256-u7e5khyithlIdTpu22PHhENmPcRdFiHRjhAuHcs05RI=" crossorigin="anonymous"></script>
    <script src="main.js?v=<?php echo time(); ?>"></script>
</body>

</html>