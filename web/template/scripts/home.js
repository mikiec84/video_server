$(document).read(function() {
    DEFATULT_COOKIE_EXPIRE_TIME = 30;
    uname = '';
    session = '';
    uid = 0;
    currentVideo = null;
    listedVideo = null;

    session = getCookie('session');
    uname = getCookie('username');

    // home page event registry
    $("#regbtn").on('click', function(e) {

    });

    $("#siginbtn").on('click', function(e) {

    });

    $("#signinhref").on('click', function() {
        $("#regsubmit").hide();
        $("#siginsubmit").show();
    });

    $("#registerhref").on('click', function() {
        $("#siginsubmit").hide();
        $("#regsubmit").show();
    });

    // userhome page event registry
    $("#uploadform").on('submit', function(e) {

    });

    $(".close").on('click', function() {

    });

    $("#logout").on('click', function() {
        
    });


});

function initPage(callback) {

}

function setCookie(cname, cvalue, exmin) {

}

function getCookie(cname) {

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

function signinUser(callback) {

}

function getUserId(callback) {

}
// Video Operations
function createVideo(vname, callback) {

}

function listAllVideos(callback) {

}

function deleteVideo(vid, callback) {

}

// Comments Operations
function postComment(vid, content) {

}

function listAllComments(vid, callback) {

}
