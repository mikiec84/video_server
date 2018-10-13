 $(function() {
    DEFATULT_COOKIE_EXPIRE_TIME = 30;
    uname = '';
    session = '';
    uid = 0;
    currentVideo = null;
    listedVideo = null;

    session = getCookie('session');
    uname = getCookie('username');
    userid = getCookie('userid');

    //home page event registry
    $("#regbtn").bind('click', function(e) {
        alert("regbtn");
    });

    $("#siginbtn").bind('click', function(e) {
        var userName = $('#userNameTxt').val()
        var password = $('#pwdTxt').val()

        if (userName.length == 0 || password.length == 0) {
            alert("User Name or password is null");
            return ;
        }
        signinUser(userName, password, getUserId);
    });

    $("#signinhref").bind('click', function() {
        alert("signinhref");
    });
    
    $("#registerhref").bind('click', function() {
        $("#siginsubmit").hide();
        $("#regsubmit").show();
        alert("registerhref");
    });

    // userhome page event registry
    $("#uploadform").bind('submit', function(e) {
        alert("uploadform");
    });

    $(".close").bind('click', function() {
        alert("close");
    });

    $("#logout").bind('click', function() {
        alert("logout");
    });

    $('#cmtBtn').bind('click', function() {

        var comment = $(cmtTxt).val();
        postComment(comment);

    });

    $('#vdoFile').bind('click', function() {
        //alert("file onload");
        $('#fileName').click();
        
    });
    $('#fileName').bind('change', function() {
       
        var filename = $(this).val();
        var array = filename.split('\\');
        filename = array[array.length - 1];
        
        $.ajaxFileUpload ({
            type: 'POST',
            url: 'http://192.168.189.134:8080/upload/' + filename,
            secureuri: false,
            fileElementId: 'fileName',
            dataType: 'json',

            success: function (data, status) {

                userid = getCookie('userid');
                session = getCookie('session');
                uname = getCookie('username');

                var req_data = {
                    "url": "http://192.168.189.134:8000/videos/wangbojing",
                    "method": "POST",
                    "req_body": "{\"author_id\": " + userid + ", \"name\":\"" + filename + "\"}"
                };

                $.ajax({
                    type: 'POST',
                    url: "http://192.168.189.134:8080/api",
                    data: JSON.stringify(req_data),
                    dataType: 'json',
                    beforeSend: function(xhrq) {
                        xhrq.setRequestHeader("X-Session-Id", session);
                        xhrq.setRequestHeader("X-User-Name", uname);
                    },
                    success: function(d) {

                    },
                    complete: function() {
                        window.location.reload();
                    },
                    error: function(e) {

                    }
                });
            },
            error: function (data, status, e) {
                alert("function error");
            }
        });

    });

    $('#logoutBtn').bind('click', function() {
        alert("logout button");
    });

});

var uname_videos;

function initPage(callback) {

}

function setCookie(cname, cvalue, exmin) {

    var exp = new Date();
    exp.setTime(exp.getTime() + exmin)

    document.cookie = cname + '=' + escape(cvalue) + ";expires=" + exp.toGMTString(); 

}

function getCookie(cname) {
    var arr, reg = new RegExp("(^| )" + cname + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg)) {
        return unescape(arr[2]);
    }
    else 
        return null;
}

//DOM Operation
function selectVideo(vid) {

}

function refreshComments(vid) {
     
}

function popupNotificationMsg(msg) {

}

function popupErrorMsg(msg) {

}

function htmlCommentListElement(cid, author, content) {

}

function htmlVideoListElement(vid, name, ctime) {

}

//Ajax
function registerUser(callback) {

}

function signinUser(userName, password, callback) {

    var hash = md5(password);
    var data = {
        "url": "http://192.168.189.134:8000/user/wangbojing",
        "method": "POST",
        "req_body": "{\"user_name\":\"" + userName + "\",\"pwd\": \"" + hash + "\"}"
    };
    
    $.ajax({
        type: 'POST',
        url: "http://192.168.189.134:8080/api",
        data: JSON.stringify(data),
        dataType: 'json',
        
        success: function(d) {
            var json = JSON.parse(JSON.stringify(d));

            setCookie("username", userName, 30 * 60 * 60);
            setCookie("session", hash, 30 * 60 * 60);

            location.href="http://192.168.189.134:8080/userHome";
        },
        error: function(e) {
            alert('error');
        },
        complete: function() {
            callback();            
        }
    });
}

function getUserId() {
    session = getCookie('session');
    uname = getCookie('username');

    var req_data = {
        "url": "http://192.168.189.134:8000/user/wangbojing",
        "method": "GET",
        "req_body": "{\"user_name\":\"" + userName + "\",\"pwd\": \"" + hash + "\"}"
    };

    $.ajax({
        type: 'POST',
        url: "http://192.168.189.134:8080/api",
        data: JSON.stringify(req_data),
        dataType: 'json',

        beforeSend: function(xhrq) {
            xhrq.setRequestHeader("X-Session-Id", session);
            xhrq.setRequestHeader("X-User-Name", uname);
        },
        success: function(d) {

            var json = JSON.parse(JSON.stringify(d));
            var idx = json.id;

            setCookie("userid", idx, 30 * 60 * 60);

        },
        error: function(e) {
            alert('error' + JSON.stringify(e));
        }
    });
}
// Video Operations
function createVideo(vname, callback) {

}

function listAllVideos(callback) {

}

function deleteVideo(vid, callback) {

}

function showDefaultVideo() {
    return video_url_base + getCookie("defaultvideo");
}

// Comments Operations
function postComment(comment) {
    userid = getCookie('userid');
    videoid = getCookie('defaultvideo_id');

    var req_data = {
        "url": "http://192.168.189.134:8000/comments/videos/" + videoid,
        "method": "POST",
        "req_body": "{\"author_id\":" + userid + ", \"content\": \"" + comment + "\"}"
    };
    
    $.ajax({
        type: 'POST',
        url: "http://192.168.189.134:8080/api",
        data: JSON.stringify(req_data),
        dataType: 'json',
        beforeSend: function(xhrq) {
            xhrq.setRequestHeader("X-Session-Id", session);
            xhrq.setRequestHeader("X-User-Name", uname);
        },
        success: function(d) {
            //alert(JSON.stringify(d));
        },
        complete: function() {
            window.location.reload();
        },
        error: function(e) {

        }
    });
}

function listAllComments(vid, callback) {

    uname = getCookie('username');
    session = getCookie('session');

    var comment_req = {
        "url": "http://192.168.189.134:8000/comments/videos/" + vid,
        "method": "GET",
        "req_body": ""
    };
    $.ajax({
        type: 'POST', 
        url: "http://192.168.189.134:8080/api",
        secureuri: false,
        dataType: 'json',
        data: JSON.stringify(comment_req),
        
        beforeSend: function(xhrq) {
            xhrq.setRequestHeader("X-Session-Id", session);
            xhrq.setRequestHeader("X-User-Name", uname);
        },
//{usericon: image_url_base + "clay_davis.png", username: 'Clay Davis', posttime: '2018-09-25 15:54:13', postcontent: 'This is the default theme that comes with SideComments.js.'},
        success: function(d) {

            var json = JSON.parse(JSON.stringify( d ));
            var comments = json.comments;

            if (comments == null) {
                AllComments = null;
                return;
            } 
            if (comments.length < 1) {
                AllComments = null;
                return ;
            } 

            AllComments = new Array(comments.length);
            for (var i = 0;i < comments.length;i ++) {
                AllComments[i] = {usericon: image_url_base + comments[i].icon,
                    username: comments[i].author,
                    posttime: comments[i].time,
                    postcontent: comments[i].content};
            }
        },
        complete: function() {
            $('.commentable-container').empty();
            $('.commentable-container').prepend(listComments(AllComments));
        }

    });

}

function uploadComplete(evt) {
    var data = JSON.parse(evt.target.responseText);
    if (data.success) {
        alert("上传成功");
    } else {
        alert("上传失败");
    }
}

var video_url_base = "http://192.168.189.134:8080/statics/video/";
var image_url_base = "http://192.168.189.134:8080/statics/images/";

var AllVideos;

function showVideoThumbnail(VideoItem, idx) {
    
    var div = '<div class="thumbnail" onclick="clickVideoThumbnail('+ idx +')">' 
            + '<div class="caption">'
            + '<p>'
            + '<video width="100%" src="'+ video_url_base+VideoItem.name +'" controls="controls">'
            + VideoItem.title
            + '</video>'
            + '</p>'
            /*
            + '<p>'
            + VideoItem.detail
            + '</p>' */
            + '</div>'
            + '</div>';
    
    return div;
}

function listVideos(Videos) {
    var html = '';
    for (var i = 0;i < Videos.length;i ++) {
        html += showVideoThumbnail(Videos[i], i);
    }
    return html;
} 
/*
var AllComments = [
    {usericon: image_url_base + "user1-128x128.jpg", username: 'Sarah Ross', posttime: '2018-09-27 15:56:34', postcontent: 'This is the default theme that comes with SideComments.js. You can easily theme and just styling it all yourself.'},
    {usericon: image_url_base + "clay_davis.png", username: 'Clay Davis', posttime: '2018-09-25 15:54:13', postcontent: 'This is the default theme that comes with SideComments.js.'},
    {usericon: image_url_base + "donald_draper.png", username: 'Donald Draper', posttime: '2018-09-24 15:56:45', postcontent: 'You can easily theme and just styling it all yourself.'},
    {usericon: image_url_base + "jon_snow.png", username: 'Jon Snow', posttime: '2018-09-21 15:56:45', postcontent: 'You can easily theme and just styling it all yourself.'},
];*/

var AllComments;

function showCommentItem(CommentItem, idx) {

    var div = '<div>'
            + '<div class=\'user-block\'>'
            + '<img class=\'img-circle img-bordered-sm\' src="' + CommentItem.usericon + '" alt=\'user image\'></img>'
            + '<span class=\'username\'>'
            + '<a href="#">' + CommentItem.username + '</a>'
            + '<a href=\'#\' class=\'pull-right btn-box-tool\'><i class=\'fa fa-times\'></i></a>'
            + '</span>'
            + '<span class=\'description\'>' + CommentItem.posttime + '</span>'
            + '</div>'
            + '<p data-section-id="' + idx + '" class="commentable-section">'
            + CommentItem.postcontent
            + '</p>'
            + '</div>'
            + '<hr style="height:1px;border:none;border-top:1px solid #555555;" />';

    return div;
}

function listComments(Comments) {

    if (Comments == null) return '';

    var html = '';
    for (var i = 0;i < Comments.length;i ++) {
        html += showCommentItem(Comments[i], i+1);
    }
    return html;
}

function clickVideoThumbnail(idx) {
    //alert('clickVideoThumbnail  ' + idx);
    //AllVideos[idx].

    setCookie("defaultvideo", AllVideos[idx].name, 30 * 60 * 60);
    setCookie("defaultvideo_id", AllVideos[idx].videoid, 30 * 60 * 60);
    setCookie("defaultvideo_authorid", AllVideos[idx].author_id, 30 * 60 * 60);

    document.getElementById('mainvideo').setAttribute("src", showDefaultVideo());

    listAllComments(AllVideos[idx].videoid, null);
}


function initUserHomePage() {

    uname = getCookie('username');
    session = getCookie('session');

    var video_req = {
        "url": "http://192.168.189.134:8000/videos/" + uname,
        "method": "GET",
        "req_body": ""
    };

    $.ajax({
        type: 'POST', 
        url: "http://192.168.189.134:8080/api",
        secureuri: false,
        dataType: 'json',
        data: JSON.stringify(video_req),
        
        beforeSend: function(xhrq) {
            xhrq.setRequestHeader("X-Session-Id", session);
            xhrq.setRequestHeader("X-User-Name", uname);
        },
        complete: function () {
            $('.col-sm-5').prepend(listVideos(AllVideos));
            //document.getElementById('col-sm-5').innerHTML = listVideos(AllVideos);
        },
        success: function(d) {
            
            var json = JSON.parse(JSON.stringify( d ));
            var videos = json.videos;

            if (videos.length < 1) return ;

            AllVideos = new Array(videos.length);
            for (var i = 0;i < videos.length;i ++) {
                AllVideos[i] = {name: videos[i].name, 
                                title: "videotag " + videos[i].name, 
                                detail: '详细内容',
                                videoid: videos[i].id,
                                author_id: videos[i].author_id};
            }

            setCookie("defaultvideo", AllVideos[0].name, 30 * 60 * 60);
            setCookie("defaultvideo_id", AllVideos[0].videoid, 30 * 60 * 60);
            setCookie("defaultvideo_authorid", AllVideos[0].author_id, 30 * 60 * 60);

            document.getElementById('mainvideo').setAttribute("src", showDefaultVideo());

            /*  */
            listAllComments(videos[0].id, null);
        },
        error: function(e) {
            alert("failed" + e);
        }
    });

    

}
