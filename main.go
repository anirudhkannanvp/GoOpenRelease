// NAME: - ANIRUDH KANNAN V P
// EMAIL:- anirudhkannanvp16@gmail.com


package main

import (
	"context"
	"fmt"
	"log"
	"errors"
	"sort"
	"os"
	"bufio"
	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
)

// LATEST VERSIONS RETURNS A SORTED SLICE WITH THE HIGHEST VERSION AS ITS FIRST ELEMENT AND THE HIGHEST VERSION OF THE SMALLER MINOR VERSIONS IN A DESCENDING ORDER
type Versions []*semver.Version


//THESE FUNCTIONS ARE A REQUIRED PRE-DEFINITON FOR SORT
func (arr Versions) Len() int {
    return len(arr)
}
func (arr Versions) Swap(i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

func (arr Versions) Less(i, j int) bool {
    return arr[i].LessThan(*arr[j])
}
// Sort sorts the given slice of Version
func DescendingSort(versions []*semver.Version) {
    sort.Sort(sort.Reverse(Versions(versions)))
}

func ReturnVersion(versionString string) (ver *semver.Version, err error){
    if versionString[0] == 'v' {
        versionString = versionString[1:]
    }
    ver = semver.New(versionString)
    err = nil
    return
}


func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run

	var versionSlice []*semver.Version

	if minVersion == nil || len(releases)==0 {
			return versionSlice
	}

	for _, release := range releases {
			if !(release.LessThan(*minVersion)) {
					versionSlice = append(versionSlice, release)
			}
	}

	if len(versionSlice) == 0 {
			return versionSlice
	}

	// SORTING IN DESCENDING OREDER
	DescendingSort(versionSlice)

	var result []*semver.Version
	result = append(result, versionSlice[0])
	var previousMaximumMajorversion, previousMaximumMinorversion = versionSlice[0].Major, versionSlice[0].Minor

	for _, release := range versionSlice {
		var currentMaximumMajorversion, currentMaximumMinorVersion = release.Major, release.Minor
			// IF CURRENT MAJOR VERSION IS LESS THAN THE PREVIOUS MAXIMUM MAJOR VERSION THEN UPDATE BOTH THE MINOR AND MAJOR VERSION
			if currentMaximumMajorversion < previousMaximumMajorversion {
					previousMaximumMinorversion, previousMaximumMajorversion = currentMaximumMinorVersion, currentMaximumMajorversion
					result = append(result, release)
					continue
			}
			// IF ONLY CURRENT MINOR VERSION IS LESS THAN THE PREVIOUS MINOR VERSION , THEN ONLY UPDATE THE PREVIOUS MINOR VERSION
			if currentMaximumMinorVersion < previousMaximumMinorversion {
					previousMaximumMinorversion = currentMaximumMinorVersion
					result = append(result, release)
			}
	}

	return result
}

func SplitString(str string) (author string, repository string, minVer *semver.Version, err error) {
    var i int
    var PositionTillNow int
		var lengthstring int
		i = 0
		lengthstring = len(str)

    // WE SET THE DEFAULT VALUES AS EMPTY AND RETURN THOSE IN CASE OF ERROR
    author, repository, minVer, err   = "", "", nil, fmt.Errorf("THE STRING IS EMPTY. ERROR")

		if lengthstring==0 {
				err = fmt.Errorf("THE STRING BECAME EMPTY BEFORE FULLY PARSING. ERROR")
				log.Fatal(err)
        return
    }


    for i < lengthstring {
        if(str[i] == '/' || str[i]==','){
            author,PositionTillNow = str[0:i], i
            break
        }
    }


    if i == lengthstring {
			err = fmt.Errorf("THE STRING BECAME EMPTY BEFORE FULLY PARSING. ERROR")
			log.Fatal(err)
      return
    }

    for i < lengthstring {
        if(str[i] == '/' || str[i]==','){
            repository,PositionTillNow = str[PositionTillNow+1:i], i
            break
        }
    }

    if i == lengthstring {
			err = fmt.Errorf("THE STRING BECAME EMPTY BEFORE FULLY PARSING. ERROR")
			log.Fatal(err)
			return
    }

    minVer, err = ReturnVersion(str[PositionTillNow+1:])

		if err != nil {
			err = fmt.Errorf("THE STRING BECAME EMPTY BEFORE FULLY PARSING. ERROR")
			log.Fatal(err)
			return
    }

		err = nil

    return author, repository, minVer, err
}

// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {


	if len(os.Args) <=1 {
		panic(errors.New("THE NUMBER OF ARGUEMENTS IS LESS. PLEASE PROVIDE THE PATH CORRECTLY"))
		return
	}

	pathOfFile := os.Args[1]

	file, err := os.Open(pathOfFile) // For read access.
	if err != nil {
		log.Fatal(err)
		return
	}

	//defer function is to make sure that the file doesnt close before the enclosing function that is the main function closes
	defer file.Close()

	// Github
	client := github.NewClient(nil)
	ctx := context.Background()

	//sc is the scanner used to read the file using bufio.NewScanner
	sc := bufio.NewScanner(file)

	for sc.Scan()!=false {
			author, repository, minVer, err := SplitString(sc.Text());
			if err != nil {
					log.Fatal(err)
					continue
			}

			var minVerIsReached = false
			var isInvalidRepository = false

			opt := &github.ListOptions{PerPage: 10}

			releases, _, err := client.Repositories.ListReleases(ctx, author, repository, opt)
			allReleases := make([]*semver.Version, len(releases))
		for !minVerIsReached {

			if err != nil {
				log.Fatal(err) // Seems like a bettter way since we have the error in log
				continue
				//panic(err) // is this really a good way?
			}
			if len(releases) == 0 {  // IF THERE ARE NO RELEASES
					break
			}

			for _, release := range releases {
				versionString := *release.TagName
				version, err := ReturnVersion(versionString)
				if err != nil{
					log.Fatal(err) // Seems like a bettter way since we have the error in log
					continue
					//panic(err) // is this really a good way?
				}
				allReleases = append(allReleases, version)
					if version.LessThan(*minVer) || version.Equal(*minVer) {
							//WE STOP QUERYING IF WE FIND A VERSION EQUAL OR LESS THAN minVer
							minVerIsReached = true
					}
			}

	}


	if isInvalidRepository {
			continue
	}

	versionSlice := LatestVersions(allReleases, minVer)
	fmt.Printf("%s %s: %s", author, repository, versionSlice)


	}

	file.Close()
}
