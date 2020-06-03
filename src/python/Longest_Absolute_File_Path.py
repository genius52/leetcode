#388
#"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"

class Solution_388:
    def lengthLongestPath(self, input: str) -> int:
        dict = {}
        cnt = 0
        max_len = 0
        lines = input.split("\n")
        for line in lines:
            depth = line.count("\t")
            cur_file_len =  len(line) - depth
            if line.find(".") >= 0:
                if depth > 0:
                    max_len = max(max_len,dict[depth-1] + cur_file_len)
                else:
                    max_len = cur_file_len
            else:
                if depth == 0:
                    dict[depth] = cur_file_len + 1
                else:
                    dict[depth] = dict[depth - 1] + cur_file_len + 1
            cnt += 1
        return max_len