//视频
const videos = ".mp4,"
// 文档
const doc = ".doc,.md,.txt,.docx"
//图片
const image = ".jpg,.jpeg,.png,.gif"

export function getFileType(ext) {
    if (!ext) return "文件夹"
    if (videos.includes(ext)) {
        return "视频"
    } else if (doc.includes(ext)) {
        return "文档"
    } else if (image.includes(ext)) {
        return "图片"
    } else {
        return "其他"
    }
}
